// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestprojectrepoaccess

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/nora-test-project-repo-access-go/internal/apijson"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/internal/requestconfig"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/option"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/packages/param"
	"github.com/stainless-sdks/nora-test-project-repo-access-go/packages/respjson"
)

// FridgeService contains methods and other services that help with interacting
// with the nora-test-project-repo-access API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFridgeService] method instead.
type FridgeService struct {
	Options []option.RequestOption
}

// NewFridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFridgeService(opts ...option.RequestOption) (r FridgeService) {
	r = FridgeService{}
	r.Options = opts
	return
}

// Add a new item to the fridge.
func (r *FridgeService) AddItem(ctx context.Context, body FridgeAddItemParams, opts ...option.RequestOption) (res *Food, err error) {
	opts = append(r.Options[:], opts...)
	path := "fridge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete an item from the fridge.
func (r *FridgeService) DeleteItem(ctx context.Context, name string, opts ...option.RequestOption) (res *FridgeDeleteItemResponse, err error) {
	opts = append(r.Options[:], opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("fridge/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all items in the fridge.
func (r *FridgeService) ListItems(ctx context.Context, opts ...option.RequestOption) (res *[]Food, err error) {
	opts = append(r.Options[:], opts...)
	path := "fridge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a specific item from the fridge.
func (r *FridgeService) GetItem(ctx context.Context, name string, opts ...option.RequestOption) (res *Food, err error) {
	opts = append(r.Options[:], opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("fridge/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing item in the fridge.
func (r *FridgeService) UpdateItem(ctx context.Context, name string, body FridgeUpdateItemParams, opts ...option.RequestOption) (res *Food, err error) {
	opts = append(r.Options[:], opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("fridge/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Model for food items in the fridge.
type Food struct {
	Name     string `json:"name,required"`
	Quantity string `json:"quantity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Food) RawJSON() string { return r.JSON.raw }
func (r *Food) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Food to a FoodParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FoodParam.Overrides()
func (r Food) ToParam() FoodParam {
	return param.Override[FoodParam](r.RawJSON())
}

// Model for food items in the fridge.
//
// The properties Name, Quantity are required.
type FoodParam struct {
	Name     string `json:"name,required"`
	Quantity string `json:"quantity,required"`
	paramObj
}

func (r FoodParam) MarshalJSON() (data []byte, err error) {
	type shadow FoodParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FoodParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FridgeDeleteItemResponse = any

type FridgeAddItemParams struct {
	// Model for food items in the fridge.
	Food FoodParam
	paramObj
}

func (r FridgeAddItemParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Food)
}
func (r *FridgeAddItemParams) UnmarshalJSON(data []byte) error {
	return r.Food.UnmarshalJSON(data)
}

type FridgeUpdateItemParams struct {
	// Model for food items in the fridge.
	Food FoodParam
	paramObj
}

func (r FridgeUpdateItemParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Food)
}
func (r *FridgeUpdateItemParams) UnmarshalJSON(data []byte) error {
	return r.Food.UnmarshalJSON(data)
}
