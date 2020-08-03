import { AxiosInstance } from "axios";
import { Token } from "../models/oauth";

type Scope = '*' | 'read.*'
type GrantTypes = 'password' | 'refresh_token'
type grandType = {
  'refresh_token': {
    refresh_token: string;
  },
  'password': {
    username: string;
    password: string;
  }
}

export class AuthAPI {
  constructor(
    private readonly axios: AxiosInstance,
    private readonly clientId: number,
    private readonly clientSecret: string
  ) {}

  public async token<T extends GrantTypes>(scope: Scope, grant_type: T, body: grandType[T]): Promise<Token> {
    const { data } = await this.axios.post(
      "/oauth/token",
      {
        client_id: this.clientId,
        client_secret: this.clientSecret,
        grant_type,
        scope,
        ...body
      }
    );

    return data;
  }

  public async loginByPwd(username: string, password: string): Promise<Token> {
    return this.token('*', 'password', {
      username,
      password
    });
  }

  public async refreshToken(refresh_token: string): Promise<Token> {
    return this.token('*', 'refresh_token', {
      refresh_token,
    });
  }
}
