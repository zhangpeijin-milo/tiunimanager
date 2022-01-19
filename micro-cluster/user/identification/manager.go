/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 * Unless required by applicable law or agreed to in writing, software        *
 * distributed under the License is distributed on an "AS IS" BASIS,          *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
 * See the License for the specific language governing permissions and        *
 * limitations under the License.                                             *
 ******************************************************************************/

package identification

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/pingcap-inc/tiem/common/constants"
	"github.com/pingcap-inc/tiem/common/errors"
	"github.com/pingcap-inc/tiem/message"
	"github.com/pingcap-inc/tiem/models"
)

type Manager struct{}

func NewIdentificationManager() *Manager {
	return &Manager{}
}

func (p *Manager) Login(ctx context.Context, request message.LoginReq) (message.LoginResp, error) {
	resp := message.LoginResp{}
	user, err := models.GetAccountReaderWriter().GetUserByName(ctx, request.Name)
	if err != nil {
		return resp, errors.NewError(errors.TIEM_LOGIN_FAILED, "incorrect username or password")
	}

	loginSuccess, err := user.CheckPassword(request.Password)
	if err != nil {
		return resp, errors.WrapError(errors.TIEM_LOGIN_FAILED, "incorrect username or password", err)
	}

	if !loginSuccess {
		return resp, errors.NewError(errors.TIEM_LOGIN_FAILED, "incorrect username or password")
	}

	// create token
	tokenString := uuid.New().String()
	expirationTime := time.Now().Add(constants.DefaultTokenValidPeriod)
	_, err = models.GetTokenReaderWriter().CreateToken(ctx, tokenString, user.ID, user.DefaultTenantID, expirationTime)
	if err != nil {
		return resp, errors.WrapError(errors.TIEM_UNRECOGNIZED_ERROR, "login failed", err)
	}

	resp.TokenString = tokenString
	resp.UserID = user.ID
	resp.TenantID = user.DefaultTenantID

	return resp, nil
}

func (p *Manager) Logout(ctx context.Context, req message.LogoutReq) (message.LogoutResp, error) {
	resp := message.LogoutResp{UserID: ""}
	token, err := models.GetTokenReaderWriter().GetToken(ctx, req.TokenString)
	if err != nil {
		return resp, errors.NewError(errors.TIEM_UNAUTHORIZED_USER, "unauthorized")
	}

	if !token.IsValid() {
		return resp, nil
	}

	resp.UserID = token.UserID
	token.Destroy()

	return resp, nil
}

func (p *Manager) Accessible(ctx context.Context, request message.AccessibleReq) (message.AccessibleResp, error) {
	resp := message.AccessibleResp{}
	token, err := models.GetTokenReaderWriter().GetToken(ctx, request.TokenString)
	if err != nil {
		return resp, err
	}

	if !token.IsValid() {
		return resp, errors.Error(errors.TIEM_ACCESS_TOKEN_EXPIRED)
	}

	resp.UserID = token.UserID
	resp.TenantID = token.TenantID

	return resp, nil
}
