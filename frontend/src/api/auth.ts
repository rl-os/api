import { AxiosInstance } from "axios";
import { Token } from "../models/oauth";
import { Config } from "../store/config";

export class AuthAPI {
  constructor(private readonly axios: AxiosInstance, private readonly config: Config) {}

  public async loginByPwd(username: string, password: string): Promise<Token> {
    const { data } = await this.axios.post(
      "/oauth/token",
      {
        client_id: this.config.clientId,
        client_secret: this.config.clientSecret,
        scope: '*', // WARNING: replace
        grant_type: 'password',
        username,
        password,
      }
    );

    return data;
  }
}
