import { DeleteIcon } from '@chakra-ui/icons';
import {
  FormControl,
  FormLabel,
  IconButton,
  Input,
  Textarea,
  VStack,
} from '@chakra-ui/react';
import { useFormContext } from 'react-hook-form';

export type JobType = {
  name: string;
  shell: string;
};

type JobProps = {
  index: number;
  job: JobType;
  remove: () => void;
};

export default function Job({ index, job, remove }: JobProps) {
  const { register } = useFormContext();

  return (
    <VStack mb={3} spacing={3}>
      <FormControl>
        <FormLabel>Name</FormLabel>
        <Input
          defaultValue={job.name}
          {...register(`jobs.${index}.name`, {
            required: true,
          })}
        />
      </FormControl>
      <FormControl>
        <FormLabel>Shell</FormLabel>
        <Textarea
          {...register(`jobs.${index}.shell`, {
            required: true,
          })}
        />
      </FormControl>
      <IconButton
        alignSelf="flex-end"
        aria-label="Delete stage"
        icon={<DeleteIcon />}
        colorScheme="red"
        onClick={remove}
      />
    </VStack>
  );
}
