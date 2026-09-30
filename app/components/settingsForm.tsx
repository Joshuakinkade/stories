"use client";

import { saveSettings } from "@/app/lib/actions/settings";
import { SaveSettingsState } from "@/app/lib/types";
import { Button, Description, FieldError, Fieldset, Form, Input, Label, Modal, TextField } from "@heroui/react";
import { useActionState, useState } from "react";

function UpdatePasswordForm() {
  const [showPasswordField, setShowPasswordField] = useState(false);
  return (<Modal>
    <Button onClick={() => setShowPasswordField(true)}>Update Password</Button>
    {showPasswordField && <Modal.Backdrop>
      <Modal.Container>
        <Modal.Dialog>
          <Modal.Header>Update Password</Modal.Header>
          <Modal.Body>
            <Form>
              <Fieldset>
                <Fieldset.Group>
                <TextField variant="secondary">
                  <Label>Current Password</Label>
                  <Input type="password" name="current_password"/>
                  <FieldError/>
                </TextField>
                <TextField variant="secondary">
                  <Label>New Password</Label>
                  <Input type="password" name="new_password"/>
                  <FieldError/>
                </TextField>
                <TextField variant="secondary">
                  <Label>Confirm Password</Label>
                  <Input type="password" name="confirm_password"/>
                  <FieldError/>
                  </TextField>
                </Fieldset.Group>
              <Fieldset.Actions>
                <Button type="submit">Save Password</Button>
                </Fieldset.Actions>
              </Fieldset>
            </Form>
          </Modal.Body>
        </Modal.Dialog>
      </Modal.Container>
    </Modal.Backdrop>}
  </Modal>
  );
}

export default function SettingsForm() {
  const [settingsState, saveSettingsAction, savePending] = useActionState<SaveSettingsState, FormData>(saveSettings, null);

  return (<Form action={saveSettingsAction}>
    <Fieldset>
      <Fieldset.Group>
        <Fieldset.Legend>Account Settings</Fieldset.Legend>
        <TextField variant="secondary">
          <Label>Display Name</Label>
          <Input name="display_name" type="text" placeholder="Josh K" />
          <Description>Will show as joshua-kinkade in urls.</Description>
          <FieldError />
        </TextField>
        <UpdatePasswordForm/>
      </Fieldset.Group>
      <Fieldset.Group>
        <Fieldset.Legend>Gallery Settings</Fieldset.Legend>
        <Description>Customize the look and feel of your public gallery.</Description>
        <TextField variant="secondary">
          <Label>Background Color</Label>
          <Input name="bg_color" type="text" placeholder="#fcf729" />
          <FieldError />
        </TextField>
      </Fieldset.Group>
      <Fieldset.Actions>
        <Button type="submit" variant="primary" isDisabled={savePending}>Save</Button>
      </Fieldset.Actions>
    </Fieldset>
  </Form>)
}
