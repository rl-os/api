import Axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import { Store } from "../store";
import { AuthAPI } from "./auth";

export class Api {
  protected axios: AxiosInstance;

  public readonly auth: AuthAPI;

  constructor(private readonly store: Store, config?: AxiosRequestConfig) {
    this.axios = Axios.create({
      ...config,
      baseURL: this.store.config.baseAPI,
    });

    this.axios.interceptors.request.use(
      this.setAccessToken,
      err => err,
    );

    this.auth = new AuthAPI(this.axios, this.store.config);
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
