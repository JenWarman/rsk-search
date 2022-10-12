package data

import (
	"encoding/json"
	"facette.io/natsort"
	"fmt"
	"github.com/pkg/errors"
	"github.com/warmans/rsk-search/pkg/models"
	"github.com/warmans/rsk-search/pkg/util"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
)

func ReplaceEpisodeFile(dataDir string, ep *models.Transcript) error {
	return util.WithReplaceJSONFileEncoder(path.Join(dataDir, fmt.Sprintf("%s.json", models.EpIDFromTranscript(ep))), func(encoder *json.Encoder) error {
		return encoder.Encode(ep)
	})
}

func EpisodeFileName(dataDir string, ep *models.Transcript) string {
	return path.Join(dataDir, fmt.Sprintf("%s.json", models.EpIDFromTranscript(ep)))
}

func SaveEpisodeToFile(dataDir string, ep *models.Transcript) error {
	return util.WithCreateJSONFileEncoder(EpisodeFileName(dataDir, ep), func(encoder *json.Encoder) error {
		return encoder.Encode(ep)
	})
}

func LoadEpisodeFile(dataDir string, fullName string) (*os.File, error) {
	return os.Open(path.Join(dataDir, fmt.Sprintf("%s.json", fullName)))
}

func LoadEpisdeByEpisodeID(dataDir string, epID string) (*models.Transcript, error) {
	f, err := LoadEpisodeFile(dataDir, epID)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	e := &models.Transcript{}

	dec := json.NewDecoder(f)
	return e, dec.Decode(e)
}

func LoadEpisodeByName(dataDir string, publication string, name string) (*models.Transcript, error) {
	return LoadEpisdeByEpisodeID(dataDir, fmt.Sprintf("ep-%s-%s", publication, name))
}

func LoadEpisodeByShortID(dataDir string, shortId string) (*models.Transcript, error) {
	return LoadEpisdeByEpisodeID(dataDir, fmt.Sprintf("ep-%s", shortId))
}

func LoadEpisodePath(path string) (*models.Transcript, error) {
	episode := &models.Transcript{}
	if err := util.WithReadJSONFileDecoder(path, func(dec *json.Decoder) error {
		return dec.Decode(episode)
	}); err != nil {
		return nil, err
	}
	return episode, nil
}

func LoadAllEpisodes(dataDir string) ([]*models.Transcript, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}
	episodes := []*models.Transcript{}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		episodePath := path.Join(dataDir, entry.Name())
		ep, err := LoadEpisodePath(episodePath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load path %s", episodePath)
		}
		episodes = append(episodes, ep)
	}
	sort.Slice(episodes, func(i, j int) bool {
		return natsort.Compare(episodes[i].ID(), episodes[j].ID())
	})
	return episodes, nil
}

func NewEpisodeStore(dataDir string) (*EpisodeCache, error) {

	episodes, err := LoadAllEpisodes(dataDir)
	if err != nil {
		return nil, err
	}

	store := &EpisodeCache{
		episodeList: make([]models.Transcript, len(episodes), len(episodes)),
		episodeMap:  map[string]models.Transcript{},
		lock:        sync.RWMutex{},
	}
	store.lock.Lock()
	defer store.lock.Unlock()
	for k, ep := range episodes {
		store.episodeMap[ep.ID()] = *ep
		store.episodeList[k] = *ep
	}
	return store, nil
}

var ErrNotFound = errors.New("not found")

type EpisodeCache struct {
	episodeMap  map[string]models.Transcript
	episodeList []models.Transcript
	lock        sync.RWMutex
}

func (s *EpisodeCache) GetEpisode(id string) (*models.Transcript, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	ep, ok := s.episodeMap[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &ep, nil
}

func (s *EpisodeCache) ListEpisodes() []*models.Transcript {

	// copy the episodeList when fetched to avoid unexpected modifications.
	list := make([]*models.Transcript, len(s.episodeList), len(s.episodeList))

	s.lock.RLock()
	defer s.lock.RUnlock()
	for k, v := range s.episodeList {
		list[k] = transcriptP(v)
	}
	return list
}

func transcriptP(transcript models.Transcript) *models.Transcript {
	return &transcript
}
