export interface DetailedUser {
  avatarURL:                        string;
  countryCode:                      string;
  defaultGroup:                     string;
  id:                               number;
  isActive:                         boolean;
  isBot:                            boolean;
  isOnline:                         boolean;
  isSupporter:                      boolean;
  lastVisit:                        Date;
  pmFriendsOnly:                    boolean;
  profileColour:                    null;
  username:                         string;
  coverURL:                         string;
  discord:                          string;
  hasSupported:                     boolean;
  interests:                        null;
  joinDate:                         Date;
  kudosu:                           Kudosu;
  lastfm:                           null;
  location:                         null;
  maxBlocks:                        number;
  maxFriends:                       number;
  occupation:                       string;
  playmode:                         string;
  playstyle:                        string[];
  postCount:                        number;
  profileOrder:                     string[];
  skype:                            null;
  title:                            null;
  twitter:                          string;
  website:                          string;
  country:                          Country;
  cover:                            Cover;
  accountHistory:                   any[];
  activeTournamentBanner:           any[];
  badges:                           any[];
  favouriteBeatmapsetCount:         number;
  followerCount:                    number;
  graveyardBeatmapsetCount:         number;
  groups:                           any[];
  lovedBeatmapsetCount:             number;
  monthlyPlaycounts:                MonthlyPlaycount[];
  page:                             Page;
  previousUsernames:                any[];
  rankedAndApprovedBeatmapsetCount: number;
  replaysWatchedCounts:             any[];
  scoresFirstCount:                 number;
  statistics:                       Statistics;
  supportLevel:                     number;
  unrankedBeatmapsetCount:          number;
  userAchievements:                 UserAchievement[];
  rankHistory:                      RankHistory;
}

export interface Country {
  code: string;
  name: string;
}

export interface Cover {
  customURL: null;
  url:       string;
  id:        string;
}

export interface Kudosu {
  total:     number;
  available: number;
}

export interface MonthlyPlaycount {
  startDate: Date;
  count:     number;
}

export interface Page {
  html: string;
  raw:  string;
}

export interface RankHistory {
  mode: string;
  data: number[];
}

export interface Statistics {
  level:                  Level;
  pp:                     number;
  ppRank:                 number;
  rankedScore:            number;
  hitAccuracy:            number;
  playCount:              number;
  playTime:               number;
  totalScore:             number;
  totalHits:              number;
  maximumCombo:           number;
  replaysWatchedByOthers: number;
  isRanked:               boolean;
  gradeCounts:            GradeCounts;
  rank:                   Rank;
}

export interface GradeCounts {
  ss:  number;
  ssh: number;
  s:   number;
  sh:  number;
  a:   number;
}

export interface Level {
  current:  number;
  progress: number;
}

export interface Rank {
  global:  number;
  country: number;
}

export interface UserAchievement {
  achievedAt:    Date;
  achievementID: number;
}
