import { useToast } from "@chakra-ui/react";
import { ToastStatusType } from "../constant/constant";

interface MessageToastProps {
  status: ToastStatusType;
  description: string;
}
function MessageToast() {
  const toast = useToast();
  const addToast = (newRes: MessageToastProps) => {
    toast({
      description: newRes.description,
      status: newRes.status,
      duration: 3000,
      isClosable: true,
      variant: "left-accent",
      position: "top-right",
    });
  };
  return { addToast };
}

export default MessageToast;
