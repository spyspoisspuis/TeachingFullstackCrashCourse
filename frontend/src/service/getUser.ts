import axios from "axios";
import { DisplayUser } from "../models/user";

export const getUserService = async () => {
  try {
    const response = await axios.get("http://localhost:8080/users/");
    if (response.data.user) {
        return [];
    }
    if (response.data.users.length === 0) {
        return [];
    }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const displayUsers: DisplayUser[] = response.data.users.map((user: any) => {
      return {
        username: user.username,
        password: user.password,
      };
    });

    return displayUsers;
  } catch {
    throw new Error("เกิดข้อผิดพลาด");
  }
};
