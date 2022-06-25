import { DeleteIcon } from '@chakra-ui/icons';
import {
  FormControl,
  FormLabel,
  HStack,
  IconButton,
  Input,
  Select,
} from '@chakra-ui/react';
import { useFormContext } from 'react-hook-form';

export type VariableType = {
  name: string;
  defaultValue: string;
  type: string;
};

type VariableProps = {
  index: number;
  variable: VariableType;
  remove: () => void;
};

export default function Variable({ variable, remove, index }: VariableProps) {
  const { register } = useFormContext();

  return (
    <HStack mb={3} align="flex-end">
      <FormControl>
        <FormLabel>Name</FormLabel>
        <Input
          defaultValue={variable.name}
          {...register(`variables.${index}.name`, {
            required: true,
          })}
        />
      </FormControl>
      <FormControl>
        <FormLabel>Type</FormLabel>
        <Select
          defaultValue={variable.type}
          {...register(`variables.${index}.type`, {
            required: true,
          })}
        >
          <option value="string">String</option>
          <option value="select">Select</option>
        </Select>
      </FormControl>
      <FormControl>
        <FormLabel>Default value</FormLabel>
        <Input
          defaultValue={variable.defaultValue}
          {...register(`variables.${index}.defaultValue`, {
            required: true,
          })}
        />
      </FormControl>
      <IconButton
        aria-label="Delete variable"
        icon={<DeleteIcon />}
        colorScheme="red"
        onClick={remove}
      />
    </HStack>
  );
}
