"use client";

import { login } from "@/app/lib/actions/auth";
import { LoginState } from "@/app/lib/types";
import { Button, ErrorMessage, FieldError, FieldGroup, Fieldset, Form, Input, Label, TextField } from "@heroui/react";
import { useActionState } from "react";

export default function LoginForm() {
  const [loginState, loginAction, loginPending] = useActionState<LoginState, FormData>(login, null)
  return (
    <Form action={loginAction}>
      <Fieldset>
        <FieldGroup>
          <TextField isRequired variant="secondary">
            <Label>Email</Label>
            <Input name="email" type="email" placeholder="test@example.com" />
            <FieldError/>
          </TextField>
          <TextField isRequired variant="secondary">
            <Label>Password</Label>
            <Input name="password" type="password" placeholder="********" />
            <FieldError/>
          </TextField>
        </FieldGroup>
        {loginState?.error && <ErrorMessage>{loginState.error}</ErrorMessage>}
        <Fieldset.Actions>
          <Button type="submit" isDisabled={loginPending}>Login</Button>
        </Fieldset.Actions>
      </Fieldset>
    </Form>
  )
}
