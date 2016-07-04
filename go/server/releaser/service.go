package releaser

import "golang.org/x/net/context"

type ReleaseService struct{}

func (r *ReleaseService) GetChannelReleases(ctx context.Context, in *ChannelReleaseRequest) (*ChannelReleases, error) {
	return &ChannelReleases{
		Channel: &Channel{ID: "abc-def", Name: "Some channel"},
		Releases: []*Release{
			&Release{ID: "rel1", SHA: "abababababa"},
			&Release{ID: "rel2", SHA: "cdcdcdcdcd"},
		},
	}, nil
}

func (r *ReleaseService) GetRelease(ctx context.Context, in *ReleaseRequest) (*Release, error) {
	return &Release{ID: "rel1", SHA: "abababababa"}, nil
}
