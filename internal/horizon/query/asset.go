package query

import (
	"fmt"
	"net/url"

	"github.com/tokend/erc20-withdraw-svc/internal/horizon/page"
)

type AssetFilters struct {
	Owner  *string
	Policy *uint32
	State  *uint32
}

type AssetIncludes struct {
	Owner bool
}

func (p AssetIncludes) Prepare() url.Values {
	result := url.Values{}
	p.prepare(&result)
	return result
}

type AssetParams struct {
	Includes   AssetIncludes
	Filters    AssetFilters
	PageParams page.Params
}

func (p AssetParams) Prepare() url.Values {
	result := url.Values{}
	p.Filters.prepare(&result)
	p.PageParams.Prepare(&result)
	p.Includes.prepare(&result)
	return result
}

func (p AssetFilters) prepare(result *url.Values) {
	if p.Policy != nil {
		result.Add("filter[policy]", fmt.Sprintf("%d", *p.Policy))
	}
	if p.Owner != nil {
		result.Add("filter[owner]", fmt.Sprintf("%s", *p.Owner))
	}
	if p.State != nil {
		result.Add("filter[state]", fmt.Sprintf("%d", *p.State))
	}
}

func (p AssetIncludes) prepare(result *url.Values) {
	if p.Owner {
		result.Add("include", "owner")
	}
}

func AssetByID(code string) string {
	return fmt.Sprintf("/v3/assets/%s", code)
}

func AssetList() string {
	return "/v3/assets"
}
