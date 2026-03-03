export type AuthData = {
  accessToken: string | undefined;
  expiredAt: Date | undefined;
};

export type AuthDataGet = () => AuthData | null | undefined;
export type AuthDataUpdate = () => Promise<void>;

export type AuthOptionsSetup = {
  getAuthData: AuthDataGet;
  updateAuthData: AuthDataUpdate;
};

export type RpcParams = {
  method: string;
  params?: object;
};
