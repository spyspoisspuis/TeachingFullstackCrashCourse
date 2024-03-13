import { VStack, Box, Text } from "@chakra-ui/react";
import ParentPage from "./ParentPage";
import { useEffect, useState } from "react";
import { DisplayUser } from "../models/user";
import { getUserService } from "../service/getUser";
import { MessageToast } from "../components";
import { ToastStatus } from "../constant/constant";
function CheckUsersPage() {
  const { addToast } = MessageToast();
  const [users, setusers] = useState<DisplayUser[]>();
  useEffect(() => {
    const fetchUsers = async () => {
      await getUserService()
        .then((response): void => {
          setusers(response);
          console.log("resposne user", response);
          addToast({
            status: ToastStatus.SUCCESS,
            description: "เรียกข้อมูลผู้ใช้สำเร็จ",
          });
        })
        .catch(() => {
          addToast({
            status: ToastStatus.ERROR,
            description: "เกิดข้อผิดพลาด",
          });
        });
    };

    fetchUsers();
  }, []);
  return (
    <ParentPage>
      <VStack>
        {users?.map((user, index) => {
          return (
            <Box key={index} color="white">
              <Text>Username: {user.username}</Text>
              <Text>Password: {user.password}</Text>
            </Box>
          );
        })}
      </VStack>
    </ParentPage>
  );
}

export default CheckUsersPage;
