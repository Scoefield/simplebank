package api

import (
	mockdb "GitCode/simplebank/db/mock"
	db "GitCode/simplebank/db/sqlc"
	"GitCode/simplebank/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)


func TestGetAccount(t *testing.T) {
	account := randomAccount()

	testCases := []struct{
		name string
		accountID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "NotFound",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Accounts{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},	
		},
		{
			name: "IntrernalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Accounts{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},	
		},
		{
			name: "InvalidID",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// require.Equal(t, http.StatusBadRequest, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},	
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
		
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			// bulid stubs
			// store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
		
			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()
		
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
		
			server.router.ServeHTTP(recorder, request)
			// check resoponse
			// require.Equal(t, http.StatusOK, recorder.Code)
			// check body
			// requireBodyMatchAccount(t, recorder.Body, account)
			tc.checkResponse(t, recorder)
		})
	}
	
}

func randomAccount() db.Accounts {
	return db.Accounts{
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Accounts) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getAccount db.Accounts
	err = json.Unmarshal(data, &getAccount)
	require.NoError(t, err)
	require.Equal(t, account, getAccount)
}
