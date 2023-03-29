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
  id: Scalars['ID'];
  itemId: Scalars['Int'];
  itemPriceType: ItemPriceType;
  onsaleType: OnsaleType;
  serviceFee: Scalars['Float'];
  total: Scalars['Float'];
  userId: Scalars['ID'];
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

export type FollowParam = {
  followerId: Scalars['Int'];
  followerName: Scalars['String'];
  followingId: Scalars['Int'];
  followingName: Scalars['String'];
  status: Scalars['Boolean'];
};

export type Item = {
  __typename?: 'Item';
  collectionId?: Maybe<Scalars['Int']>;
  createDate?: Maybe<Scalars['Date']>;
  creator: Scalars['String'];
  creatorId: Scalars['ID'];
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
  onsaleType: OnsaleType;
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
  createCollection: Scalars['Boolean'];
  follow?: Maybe<Scalars['String']>;
  login: Scalars['String'];
  newUser: Scalars['String'];
  placeBid: Scalars['Boolean'];
  refreshToken: Scalars['String'];
  setPriceAndMint: Scalars['Boolean'];
  updateUser: Scalars['String'];
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


export type MutationFollowArgs = {
  param?: InputMaybe<FollowParam>;
};


export type MutationLoginArgs = {
  input: Login;
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


export type MutationSetPriceAndMintArgs = {
  param?: InputMaybe<PriceParam>;
};


export type MutationUpdateUserArgs = {
  user?: InputMaybe<UpdateUser>;
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
  itemId: Scalars['ID'];
  onsaleType: OnsaleType;
  payAmount: Scalars['Float'];
  serviceFee: Scalars['Float'];
};

export type Payment = {
  __typename?: 'Payment';
  createDate?: Maybe<Scalars['Date']>;
  createor: Scalars['String'];
  id: Scalars['ID'];
  itemId: Scalars['Int'];
  onsaleType: OnsaleType;
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
  onsaleType: OnsaleType;
  startDate?: InputMaybe<Scalars['Date']>;
};

export type PriceRange = {
  max: Scalars['Float'];
  min: Scalars['Float'];
};

export type Query = {
  __typename?: 'Query';
  collection?: Maybe<Array<Maybe<Collection>>>;
  item?: Maybe<Item>;
  items?: Maybe<Array<Maybe<Item>>>;
  searchItems?: Maybe<Array<Maybe<Item>>>;
  user?: Maybe<User>;
};


export type QueryCollectionArgs = {
  creator: Scalars['String'];
};


export type QueryItemArgs = {
  id: Scalars['ID'];
};


export type QueryItemsArgs = {
  createor: Scalars['String'];
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

export type UpdateUser = {
  avatar?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  company?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['String']>;
  img?: InputMaybe<Scalars['String']>;
  links?: InputMaybe<Scalars['String']>;
  password?: InputMaybe<Scalars['String']>;
  phone?: InputMaybe<Scalars['String']>;
  realname: Scalars['String'];
  verifyName?: InputMaybe<Scalars['String']>;
  verifyType?: InputMaybe<VerifyType>;
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
  password?: Maybe<Scalars['String']>;
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
  createId: Scalars['ID'];
  creator: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  saleStatus: Scalars['Int'];
  tag?: InputMaybe<Scalars['String']>;
  uploadUrl: Scalars['String'];
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
