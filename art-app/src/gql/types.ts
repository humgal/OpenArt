export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Date: any;
};

export type Bid = {
  __typename?: 'Bid';
  bidDate: Scalars['Date'];
  bidEndDate?: Maybe<Scalars['Date']>;
  itemPriceType: ItemPriceType;
  serviceFee: Scalars['Float'];
  temId: Scalars['Int'];
  total: Scalars['Float'];
  user: Scalars['ID'];
  username: Scalars['String'];
};

export type BidParm = {
  itemId: Scalars['Int'];
  total: Scalars['Float'];
};

export enum Blockchain {
  Bnb = 'BNB',
  Ethereum = 'ETHEREUM',
  Klaytn = 'KLAYTN',
  Matic = 'MATIC',
  Solana = 'SOLANA'
}

export type Collection = {
  __typename?: 'Collection';
  createDate?: Maybe<Scalars['Date']>;
  createor?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  items?: Maybe<Array<Maybe<Item>>>;
  name: Scalars['String'];
};

export type CollectionParm = {
  creator: Scalars['String'];
  name?: InputMaybe<Scalars['String']>;
};

export type Follow = {
  __typename?: 'Follow';
  followerId: Scalars['Int'];
  followerName: Scalars['String'];
  followingId: Scalars['Int'];
  followingName: Scalars['String'];
  id: Scalars['ID'];
  status: Scalars['Boolean'];
};

export type Item = {
  __typename?: 'Item';
  createDate?: Maybe<Scalars['Date']>;
  createorId: Scalars['ID'];
  creator: Scalars['String'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  name: Scalars['String'];
  price?: Maybe<ItemPrice>;
  saleStatus: Scalars['Int'];
  tag?: Maybe<Scalars['String']>;
  uploadUrl: Scalars['String'];
};

export type ItemPrice = {
  __typename?: 'ItemPrice';
  expirationDate?: Maybe<Scalars['Date']>;
  initPrice: Scalars['Float'];
  itemId: Scalars['ID'];
  itemPriceType: ItemPriceType;
  onsale: OnsaleType;
  serviceFee: Scalars['Float'];
  startDate?: Maybe<Scalars['Date']>;
};

export enum ItemPriceType {
  Auction = 'AUCTION',
  Fixed = 'FIXED'
}

export type Link = {
  __typename?: 'Link';
  type: LinkType;
  url: Scalars['String'];
};

export enum LinkType {
  Discord = 'DISCORD',
  Facebook = 'FACEBOOK',
  Instagram = 'INSTAGRAM',
  Tiktok = 'TIKTOK',
  Website = 'WEBSITE',
  Youtube = 'YOUTUBE'
}

export type Login = {
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  checkout?: Maybe<Scalars['String']>;
  connectWallet: Wallet;
  createCollection?: Maybe<Collection>;
  login: Scalars['String'];
  mintArt: Array<Item>;
  newUser: Scalars['String'];
  placeBid?: Maybe<Bid>;
  refreshToken: Scalars['String'];
  setPrice: Array<Item>;
  uploadArt: Array<Item>;
};


export type MutationCheckoutArgs = {
  param?: InputMaybe<PayParam>;
};


export type MutationConnectWalletArgs = {
  type: WalletType;
  userId: Scalars['ID'];
};


export type MutationCreateCollectionArgs = {
  param: CollectionParm;
};


export type MutationLoginArgs = {
  input: Login;
};


export type MutationMintArtArgs = {
  items?: InputMaybe<Array<Scalars['ID']>>;
};


export type MutationNewUserArgs = {
  user?: InputMaybe<NewUser>;
};


export type MutationPlaceBidArgs = {
  bid?: InputMaybe<BidParm>;
};


export type MutationRefreshTokenArgs = {
  input: RefreshTokenInput;
};


export type MutationSetPriceArgs = {
  param?: InputMaybe<PriceParam>;
};


export type MutationUploadArtArgs = {
  items?: InputMaybe<Array<UploadItem>>;
};

export type NewUser = {
  email?: InputMaybe<Scalars['String']>;
  password: Scalars['String'];
  phone?: InputMaybe<Scalars['String']>;
  username: Scalars['String'];
};

export enum OnsaleType {
  Eth = 'ETH',
  Weth = 'WETH',
  OxBtc = 'oxBTC'
}

export type PayParam = {
  balance: Scalars['Float'];
  id: Scalars['ID'];
  onsaleType: OnsaleType;
  payAmount: Scalars['Float'];
  serviceFee: Scalars['Float'];
};

export type Payment = {
  __typename?: 'Payment';
  createDate?: Maybe<Scalars['Date']>;
  createor: Scalars['String'];
  itemId: Scalars['Int'];
  onsaleType: Scalars['Int'];
  payDate?: Maybe<Scalars['Date']>;
  payStatus: Scalars['Int'];
  payUser?: Maybe<Scalars['String']>;
  price: Scalars['Float'];
  serviceFee: Scalars['Float'];
};

export type PriceParam = {
  expirationDate?: InputMaybe<Scalars['Date']>;
  initPrice: Scalars['Float'];
  itemId: Scalars['ID'];
  itemPriceType: ItemPriceType;
  onsale: OnsaleType;
  startDate?: InputMaybe<Scalars['Date']>;
};

export type PriceRange = {
  max: Scalars['Float'];
  min: Scalars['Float'];
};

export type Query = {
  __typename?: 'Query';
  collection?: Maybe<Collection>;
  item?: Maybe<Item>;
  items?: Maybe<Array<Maybe<Item>>>;
  searchItems?: Maybe<Array<Maybe<Item>>>;
  user?: Maybe<User>;
};


export type QueryCollectionArgs = {
  createor: Scalars['String'];
};


export type QueryItemArgs = {
  id: Scalars['ID'];
};


export type QueryItemsArgs = {
  createor?: InputMaybe<Scalars['String']>;
  ids?: InputMaybe<Array<Scalars['ID']>>;
};


export type QuerySearchItemsArgs = {
  param: SearchParm;
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};

export type RefreshTokenInput = {
  token: Scalars['String'];
};

export type SearchParm = {
  chain?: InputMaybe<Blockchain>;
  creator?: InputMaybe<Scalars['String']>;
  onsale?: InputMaybe<OnsaleType>;
  param: Scalars['String'];
  price?: InputMaybe<PriceRange>;
  type?: InputMaybe<SearchType>;
};

export enum SearchType {
  All = 'ALL',
  Animation = 'ANIMATION',
  Game = 'GAME',
  Photogrphy = 'PHOTOGRPHY',
  Video = 'VIDEO'
}

export type Subscription = {
  __typename?: 'Subscription';
  subscriptionPayment?: Maybe<SubscriptionEvent>;
};


export type SubscriptionSubscriptionPaymentArgs = {
  itemid?: InputMaybe<Scalars['String']>;
};

export type SubscriptionEvent = {
  __typename?: 'SubscriptionEvent';
  bids?: Maybe<Array<Maybe<Bid>>>;
  payments?: Maybe<Array<Maybe<Payment>>>;
};

export type User = {
  __typename?: 'User';
  avatar?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  company?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
  followerNum?: Maybe<Scalars['Int']>;
  followingNum?: Maybe<Scalars['Int']>;
  id: Scalars['ID'];
  img?: Maybe<Scalars['String']>;
  isCreator?: Maybe<Scalars['Boolean']>;
  joinDate?: Maybe<Scalars['Date']>;
  links?: Maybe<Array<Maybe<Link>>>;
  phone?: Maybe<Scalars['String']>;
  realname: Scalars['String'];
  username: Scalars['String'];
  verifyName?: Maybe<Scalars['String']>;
  verifyType?: Maybe<VerifyType>;
};

export enum VerifyType {
  Instagram = 'INSTAGRAM',
  Twitter = 'TWITTER'
}

export type UploadItem = {
  collection?: InputMaybe<Scalars['String']>;
  creator: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  saleStatus: Scalars['Int'];
  tag?: InputMaybe<Scalars['String']>;
  uploadUrl: Array<Scalars['String']>;
};

export type Wallet = {
  __typename?: 'wallet';
  pubToken: Scalars['String'];
  type: WalletType;
};

export enum WalletType {
  Bank = 'BANK',
  Coin = 'COIN',
  Wallet = 'WALLET'
}
