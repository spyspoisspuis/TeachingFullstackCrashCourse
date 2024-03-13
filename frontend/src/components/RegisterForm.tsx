import { Box, VStack, Button } from "@chakra-ui/react";
import InputField from "./InputField";
import { useState } from "react";

interface RegisterFormProps {
  register: (username: string, password: string) => void;
}

function RegisterForm({ register }: RegisterFormProps) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const submit = () => {
    if (password !== confirmPassword) {
      console.log("รหัสผ่านไม่ตรงกัน");
      return;
    }
    console.log(username, password, confirmPassword);
    register(username, password);
  };

  return (
    <Box bg="white" w="lg">
      <VStack p={8}>
        <InputField label={"ชื่อผู้ใช้งาน"} setState={setUsername} />
        <InputField label={"รหัสผ่าน"} setState={setPassword} />
        <InputField label={"ยืนยันรหัสผ่าน"} setState={setConfirmPassword} />
        <Button colorScheme="blue" onClick={submit}>
          Register
        </Button>
      </VStack>
    </Box>
  );
}

export default RegisterForm;
