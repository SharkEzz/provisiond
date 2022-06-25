import { Link, VStack } from '@chakra-ui/react';
import { NavLink } from 'react-router-dom';

export default function Sidebar() {
  return (
    <VStack as="nav" borderRightWidth={1} p={4} align="flex-start">
      <Link as={NavLink} to="/">
        Dashboard
      </Link>
    </VStack>
  );
}
