/* tslint:disable */
import {
  RskAudioQuality,
  RskMediaType,
  RskSynopsis,
} from '.';

export interface RskShortTranscript {
  actors?: string[];
  audioQuality?: RskAudioQuality;
  audioUri?: string;
  bestof?: boolean;
  episode?: number;
  id?: string;
  incomplete?: boolean;
  mediaType?: RskMediaType;
  metadata?: { [key: string]: string };
  name?: string;
  offsetAccuracyPcnt?: number;
  publication?: string;
  releaseDate?: string;
  series?: number;
  shortId?: string;
  special?: boolean;
  summary?: string;
  synopsis?: RskSynopsis[];
  transcriptAvailable?: boolean;
  triviaAvailable?: boolean;
  version?: string;
}
