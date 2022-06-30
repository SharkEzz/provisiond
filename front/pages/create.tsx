import {
  Box,
  Button,
  Divider,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Text,
} from '@chakra-ui/react';
import { FormProvider, useFieldArray, useForm } from 'react-hook-form';
import Job, { JobType } from '../components/DeploymentForm/Job';
import Variable, { VariableType } from '../components/DeploymentForm/Variable';
import Layout from '../components/Layout';

export default function CreatePage() {
  const methods = useForm<{
    variables: VariableType[];
    name: string;
    jobs: JobType[];
  }>();

  const {
    fields: variablesFields,
    append: appendVariable,
    remove: removeVariable,
  } = useFieldArray({
    control: methods.control,
    name: 'variables',
  });

  const {
    fields: jobsFields,
    append: appendJob,
    remove: removeJob,
  } = useFieldArray({
    control: methods.control,
    name: 'jobs',
  });

  const onSubmit = (data: unknown) => console.log(data);

  return (
    <Layout title="New deployment">
      {/* eslint-disable-next-line react/jsx-props-no-spreading */}
      <FormProvider {...methods}>
        <Flex
          as="form"
          gap={6}
          flexDir="column"
          onSubmit={methods.handleSubmit(onSubmit)}
        >
          <FormControl>
            <FormLabel fontSize="xl">Name</FormLabel>
            <Divider my={4} />
            <Input
              {...methods.register('name', {
                required: true,
              })}
            />
          </FormControl>
          <Box>
            <Text mb={3} fontSize="xl">
              Variables
            </Text>
            <Divider my={4} />
            {variablesFields.map((variable, index) => (
              <Variable
                index={index}
                key={variable.id}
                variable={variable}
                remove={() => removeVariable(index)}
              />
            ))}
            <Button
              onClick={() => {
                appendVariable({ name: '', type: 'string', defaultValue: '' });
              }}
            >
              Add variable
            </Button>
          </Box>
          <Box>
            <Text mb={3} fontSize="xl">
              Jobs
            </Text>
            <Divider my={4} />
            {jobsFields.map((job, index) => (
              <Job
                key={job.id}
                index={index}
                job={job}
                remove={() => removeJob(index)}
              />
            ))}
            <Button
              onClick={() => {
                appendJob({ name: '', shell: '' });
              }}
            >
              Add job
            </Button>
          </Box>
          <Button alignSelf="flex-end" type="submit" colorScheme="green">
            Create
          </Button>
        </Flex>
      </FormProvider>
    </Layout>
  );
}
