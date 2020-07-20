import { AxiosInstance } from "axios";
import { DetailedUser } from "../models/detailed_user";

export class UserAPI {
  constructor(private readonly axios: AxiosInstance) {}

  public async me(): Promise<DetailedUser> {
    const { data } = await this.axios.get(
      '/api/v2/me/'
    );

    return data;
  }
}
