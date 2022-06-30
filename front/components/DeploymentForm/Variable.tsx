import { DeleteIcon } from '@chakra-ui/icons';
import {
  FormControl,
  FormLabel,
  HStack,
  IconButton,
  Input,
  Select,
  Textarea,
} from '@chakra-ui/react';
import { useState } from 'react';
import { useFormContext } from 'react-hook-form';

export type VariableType = {
  name: string;
  defaultValue: string;
  type: 'string' | 'select';
};

type VariableProps = {
  index: number;
  variable: VariableType;
  remove: () => void;
};

export default function Variable({ variable, remove, index }: VariableProps) {
  const { register, resetField } = useFormContext();

  const [variableType, setVariableType] = useState<'string' | 'select'>(
    variable.type,
  );

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
          onChange={(e) => {
            resetField(`variables.${index}.defaultValue`);
            setVariableType(e.target.value as 'string' | 'select');
          }}
        >
          <option value="string">String</option>
          <option value="select">Select</option>
        </Select>
      </FormControl>
      <FormControl>
        <FormLabel>
          Default value {variableType === 'select' ? '(one per line)' : ''}
        </FormLabel>
        {variableType === 'string' ? (
          <Input
            defaultValue={variable.defaultValue}
            {...register(`variables.${index}.defaultValue`, {
              required: true,
            })}
          />
        ) : (
          <Textarea
            defaultValue={variable.defaultValue}
            {...register(`variables.${index}.defaultValue`, {
              required: true,
            })}
          />
        )}
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
