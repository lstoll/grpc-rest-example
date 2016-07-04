package releaser

import "golang.org/x/net/context"

type ReleaseService struct {
	Channel  *Channel
	Releases []*Release
}

func NewService() *ReleaseService {
	return &ReleaseService{
		Channel: &Channel{ID: "abc-def", Name: "Some channel"},
		Releases: []*Release{
			&Release{ID: "rel1", SHA: "abababababa"},
			&Release{ID: "rel2", SHA: "cdcdcdcdcd"},
		},
	}

}

func (r *ReleaseService) GetChannelReleases(ctx context.Context, in *ChannelReleaseRequest) (*ChannelReleases, error) {
	return &ChannelReleases{
		Channel:  r.Channel,
		Releases: r.Releases,
	}, nil
}

func (r *ReleaseService) GetRelease(ctx context.Context, in *ReleaseRequest) (*Release, error) {
	return &Release{ID: "rel1", SHA: "abababababa"}, nil
}

func (r *ReleaseService) AddRelease(ctx context.Context, in *Release) (*AddReleaseResponse, error) {
	r.Releases = append(r.Releases, in)
	return &AddReleaseResponse{}, nil
}
