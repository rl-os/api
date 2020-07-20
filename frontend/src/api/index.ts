import Axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import { Store } from "../store";
import { AuthAPI } from "./auth";
import { UserAPI } from "./user";

import Config from "../config";

export class Api {
  protected axios: AxiosInstance;

  public readonly auth: AuthAPI;
  public readonly user: UserAPI;

  constructor(private readonly store: Store, config?: AxiosRequestConfig) {
    this.axios = Axios.create({
      ...config,
      baseURL: Config.baseAPI,
    });

    this.axios.interceptors.request.use(
      req => this.setAccessToken(req),
      err => err,
    );

    this.auth = new AuthAPI(this.axios, Config.oauth.clientId, Config.oauth.clientSecret);
    this.user = new UserAPI(this.axios);
  }

  private setAccessToken(req: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...req,
      headers: {
        ...req.headers,
        "Authorization": `Bearer ${this.store.auth.accessToken}`,
      },
    };
  }
}
