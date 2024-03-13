import { Heading } from "@chakra-ui/react";
import { RegisterForm, MessageToast } from "../components";
import { RegisterModel } from "../models/register";
import { registerService } from "../service/register";
import { ToastStatus } from "../constant/constant";
import ParentPage from "./ParentPage";

function RegisterPage() {
  const { addToast } = MessageToast();
  const register = async (username: string, password: string) => {
    console.log("Page register", username, password);
    //TODO : Making api request to register
    const registerModel: RegisterModel = {
      username: username,
      password: password,
    };
    addToast({
      status: ToastStatus.INFO,
      description: "กำลังลงทะเบียน...",
    });

    await registerService(registerModel)
      .then(() => {
        addToast({
          status: ToastStatus.SUCCESS,
          description: "ลงทะเบียนสำเร็จ...",
        });
      })
      .catch(() => {
        addToast({
          status: ToastStatus.ERROR,
          description: "เกิดข้อผิดพลาด.",
        });
      });
  };

  return (
    <ParentPage>
      <Heading color="white">ลงทะเบียน</Heading>
      <RegisterForm register={register} />
    </ParentPage>
  );
}

export default RegisterPage;
