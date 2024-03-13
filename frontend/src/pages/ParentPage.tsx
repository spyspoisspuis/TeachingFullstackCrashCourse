import { ReactNode } from "react";
import { Flex } from "@chakra-ui/react";
interface ParentPageProps {
  children: ReactNode;
}
function ParentPage({ children }: ParentPageProps) {
  return (
    <Flex
      w="full"
      minH="100vh"
      bg="gray.700"
      justify={"flex-start"}
      align={"center"}
      //   direction={{ base: "column", lg: "row" }}
      direction="column"
      py={12}
    >
      {children}
    </Flex>
  );
}

export default ParentPage;
