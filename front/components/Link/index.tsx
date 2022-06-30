import { forwardRef } from 'react';
import NextLink from 'next/link';
import { Link as ChakraLink } from '@chakra-ui/react';

const Link = forwardRef(
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  ({ href, children }: { href: string; children: React.ReactNode }, _) => (
    <NextLink href={href} passHref>
      <ChakraLink>{children}</ChakraLink>
    </NextLink>
  ),
);
Link.displayName = 'Link';

export default Link;
