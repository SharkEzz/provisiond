import { Flex, Text } from '@chakra-ui/react';

export default function Navbar() {
  return (
    <Flex
      as="header"
      h={16}
      flexShrink={0}
      bg="green.600"
      color="white"
      align="center"
      px={8}
    >
      <Text as="span" fontWeight="bold" fontSize="xl">
        provisiond
      </Text>
    </Flex>
  );
}
