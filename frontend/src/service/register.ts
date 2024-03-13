import axios from "axios";
import { RegisterModel } from "../models/register";

export const registerService = async (model: RegisterModel) => {
  try {
    const response = await axios.post("http://localhost:8080/users/", {
      username: model.username,
      password: model.password,
    });
    return response.data;
  } catch {
    throw new Error("เกิดข้อผิดพลาด");
  }
};
