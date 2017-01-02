package snippets

import (
	"bitbucket.com/sharingmachine/kwkcli/models"
	"bitbucket.com/sharingmachine/kwkcli/rpc"
	"bitbucket.com/sharingmachine/kwkcli/config"
	"bitbucket.com/sharingmachine/rpc/src/aliasesRpc"
	"google.golang.org/grpc"
	"time"
)

type RpcService struct {
	Settings config.Settings
	client   aliasesRpc.AliasesRpcClient
	headers  *rpc.Headers
}

func New(conn *grpc.ClientConn, s config.Settings, h *rpc.Headers) Service {
	return &RpcService{Settings: s, client: aliasesRpc.NewAliasesRpcClient(conn), headers: h}
}

func (a *RpcService) Update(fullKey string, description string) (*models.Snippet, error) {
	if r, err := a.client.Update(a.headers.GetContext(), &aliasesRpc.UpdateRequest{FullKey: fullKey, Description: description}); err != nil {
		return nil, err
	} else {
		m := &models.Snippet{}
		mapAlias(r.Alias, m)
		return m, nil
	}
}

// since unix time in milliseconds
func (a *RpcService) List(username string, size int64, since int64, tags ...string) (*models.SnippetList, error) {
	if res, err := a.client.List(a.headers.GetContext(), &aliasesRpc.ListRequest{Username: username, Since: since, Size: size, Tags: tags}); err != nil {
		return nil, err
	} else {
		list := &models.SnippetList{}
		mapSnippetList(res, list)
		return list, nil
	}
}

func (a *RpcService) Get(k *models.KwkKey) (*models.SnippetList, error) {
	if res, err := a.client.Get(a.headers.GetContext(), &aliasesRpc.GetRequest{Username: k.Username, FullKey: k.FullKey}); err != nil {
		return nil, err
	} else {
		list := &models.SnippetList{}
		mapSnippetList(res, list)
		return list, nil
	}
}

func (a *RpcService) Delete(fullKey string) error {
	_, err := a.client.Delete(a.headers.GetContext(), &aliasesRpc.DeleteRequest{FullKey: fullKey})
	return err
}

func (a *RpcService) Create(uri string, path string) (*models.CreateSnippet, error) {
	// encrypt if requested
	if res, err := a.client.Create(a.headers.GetContext(), &aliasesRpc.CreateRequest{Uri: uri, FullKey: path}); err != nil {
		return nil, err
	} else {
		createAlias := &models.CreateSnippet{}
		if res.Alias != nil {
			alias := &models.Snippet{}
			mapAlias(res.Alias, alias)
			createAlias.Snippet = alias
		} else {
			createAlias.TypeMatch = &models.TypeMatch{
				Matches: []models.Match{},
			}
			for _, v := range res.TypeMatch.Matches {
				m := models.Match{
					Extension: v.Extension,
					Media:     v.Media,
					Runtime:   v.Runtime,
					Score:     v.Score,
				}
				createAlias.TypeMatch.Matches = append(createAlias.TypeMatch.Matches, m)
			}
		}
		return createAlias, nil
	}
}

func (a *RpcService) Rename(fullKey string, newFullKey string) (*models.Snippet, string, error) {
	if res, err := a.client.Rename(a.headers.GetContext(), &aliasesRpc.RenameRequest{FullKey: fullKey, NewFullKey: newFullKey}); err != nil {
		return nil, "", err
	} else {
		alias := &models.Snippet{}
		mapAlias(res.Alias, alias)
		return alias, res.OriginalFullKey, nil
	}
}

func (a *RpcService) Patch(fullKey string, target string, patch string) (*models.Snippet, error) {
	if res, err := a.client.Patch(a.headers.GetContext(), &aliasesRpc.PatchRequest{FullKey: fullKey, Target: target, Patch: patch}); err != nil {
		return nil, err
	} else {
		alias := &models.Snippet{}
		mapAlias(res.Alias, alias)
		return alias, nil
	}
}

func (a *RpcService) Clone(k *models.KwkKey, newFullKey string) (*models.Snippet, error) {
	if res, err := a.client.Clone(a.headers.GetContext(), &aliasesRpc.CloneRequest{Username: k.Username, FullKey: k.FullKey, NewFullKey: newFullKey}); err != nil {
		return nil, err
	} else {
		alias := &models.Snippet{}
		mapAlias(res.Alias, alias)
		return alias, nil
	}
}

func (a *RpcService) Tag(fullKey string, tags ...string) (*models.Snippet, error) {
	if res, err := a.client.Tag(a.headers.GetContext(), &aliasesRpc.TagRequest{FullKey: fullKey, Tags: tags}); err != nil {
		return nil, err
	} else {
		alias := &models.Snippet{}
		mapAlias(res.Alias, alias)
		return alias, nil
	}
}

func (a *RpcService) UnTag(fullKey string, tags ...string) (*models.Snippet, error) {
	if res, err := a.client.UnTag(a.headers.GetContext(), &aliasesRpc.UnTagRequest{FullKey: fullKey, Tags: tags}); err != nil {
		return nil, err
	} else {
		alias := &models.Snippet{}
		mapAlias(res.Alias, alias)
		return alias, nil
	}
}

func mapAlias(rpc *aliasesRpc.AliasResponse, model *models.Snippet) {
	model.Id = rpc.SnipId
	model.FullKey = rpc.FullKey
	model.Username = rpc.Username
	model.Key = rpc.Key
	model.Extension = rpc.Extension
	// if encrypted, decrypt
	// if checksum doesn't match then throw warning
	// check that checksum signature is valid with public key.
	model.Snip = rpc.Snip
	model.Version = rpc.SnipVersion
	model.Runtime = rpc.Runtime
	model.Tags = rpc.Tags
	model.Created = time.Unix(rpc.Created/1000, 0)
	model.Description = rpc.Description
	model.ForkedFromFullKey = rpc.ForkedFromFullKey
	model.ForkedFromVersion = rpc.ForkedFromVersion
	model.Private = rpc.Private
	model.RunCount = rpc.RunCount
	model.CloneCount = rpc.CloneCount
}

func mapSnippetList(rpc *aliasesRpc.AliasListResponse, model *models.SnippetList) {
	model.Total = rpc.Total
	model.Since = time.Unix(rpc.Since/1000, 0)
	model.Size = rpc.Size
	for _, v := range rpc.Items {
		item := &models.Snippet{}
		mapAlias(v, item)
		model.Items = append(model.Items, *item)
	}
}
