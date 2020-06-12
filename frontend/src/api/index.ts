import Axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import { Store } from "../store";
import { AuthAPI } from "./auth";
import { UserAPI } from "./user";

export class Api {
  protected axios: AxiosInstance;

  public readonly auth: AuthAPI;
  public readonly user: UserAPI;

  constructor(private readonly store: Store, config?: AxiosRequestConfig) {
    this.axios = Axios.create({
      ...config,
      baseURL: this.store.config.baseAPI,
    });

    this.axios.interceptors.request.use(
      req => this.setAccessToken(req),
      err => err,
    );

    this.auth = new AuthAPI(this.axios, this.store.config.clientId, this.store.config.clientSecret);
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
