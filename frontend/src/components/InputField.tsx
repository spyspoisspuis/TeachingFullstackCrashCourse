import { Input, Text, Flex } from "@chakra-ui/react";
interface InputFieldProps {
  label: string;
  setState: (value: string) => void;
}
function InputField({ label, setState }: InputFieldProps) {
  return (
    <Flex w="full" direction="column">
      <Text>{label}</Text>
      <Input onChange={(e) => {
        setState(e.target.value);
      }} />
    </Flex>
  );
}

export default InputField;
