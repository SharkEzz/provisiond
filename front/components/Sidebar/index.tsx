import { Link, VStack } from '@chakra-ui/react';

export default function Sidebar() {
  return (
    <VStack as="nav" borderRightWidth={1} p={4} align="flex-start">
      <Link href="/">Dashboard</Link>
    </VStack>
  );
}
