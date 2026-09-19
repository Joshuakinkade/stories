export type SignupState = {
  error?: {
    email?: string;
    display_name?: string;
    password?: string;
    code?: string,
}, data?: object } | null;

export type LoginState = {
  error?: string
} | null;
