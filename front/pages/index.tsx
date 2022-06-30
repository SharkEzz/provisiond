import { ChevronRightIcon, DeleteIcon, EditIcon } from '@chakra-ui/icons';
import {
  Badge,
  Flex,
  IconButton,
  Table,
  TableCaption,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import Layout from '../components/Layout';
import Link from '../components/Link';

export default function HomePage() {
  return (
    <Layout title="Dashboard">
      <Link href="/create">New</Link>
      <TableContainer boxShadow="md">
        <Table variant="striped">
          <TableCaption>List of deployments</TableCaption>
          <Thead>
            <Tr>
              <Th>Name</Th>
              <Th>Last execution</Th>
              <Th>Last execution time</Th>
              <Th>Actions</Th>
            </Tr>
          </Thead>
          <Tbody>
            <Tr>
              <Td>Wordpress</Td>
              <Td>
                <Badge colorScheme="green">success</Badge>
              </Td>
              <Td>05/05/2022</Td>
              <Td>
                <Flex flexWrap="wrap" gap={2}>
                  <IconButton
                    aria-label="Start deployment"
                    icon={<ChevronRightIcon />}
                    colorScheme="green"
                  />
                  <IconButton
                    aria-label="Edit deployment"
                    icon={<EditIcon />}
                    colorScheme="blue"
                  />
                  <IconButton
                    aria-label="Delete deployment"
                    icon={<DeleteIcon />}
                    colorScheme="red"
                  />
                </Flex>
              </Td>
            </Tr>
          </Tbody>
        </Table>
      </TableContainer>
    </Layout>
  );
}
