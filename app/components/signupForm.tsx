"use client";

import { useActionState } from "react";
import { Button, Description, FieldError, FieldGroup, Fieldset, Form, Input, Label, TextField } from '@heroui/react';
import { signup } from '../app/lib/actions/auth';
import { SignupState } from "@/app/lib/types";


export default function SignUpForm({ code }: { code: string }) {
  const [state, action, pending] = useActionState<SignupState, FormData>(signup, null);

  return (
    <Form action={action} validationErrors={state?.error}>
      <Fieldset>
        <Fieldset.Legend>Sign Up</Fieldset.Legend>
        <Description>Create an account</Description>
        <FieldGroup>
          <TextField name="email" isRequired>
            <Label>Email</Label>
            <Input type="email" placeholder="test@example.com" />
            <FieldError />
          </TextField>
          <TextField name="display_name" isRequired>
            <Label>Display Name</Label>
            <Input type="text" placeholder="John Smith" />
            <Description>This will appear on the site and is used to create a unique url for your stories.</Description>
            <FieldError />
          </TextField>
          <TextField name="password" isRequired>
            <Label>Password</Label>
            <Input type="password" placeholder="********"/>
           <FieldError />
          </TextField>
          <Input type="hidden" name="code" value={code} readOnly />
          <TextField name="confirm_password" isRequired>
            <Label>Retype Password</Label>
            <Input type="password" placeholder="********" />
            <FieldError />
          </TextField>
        </FieldGroup>
        <Fieldset.Actions>
          <Button type="submit" isDisabled={pending}>
            Sign Up
          </Button>
        </Fieldset.Actions>
      </Fieldset>
    </Form>
  )
}
