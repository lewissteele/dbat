import { Driver } from "./driver";

export interface Connection {
  database: string;
  driver: Driver;
  host: string;
  password: string;
  port: string;
  user: string;
}
