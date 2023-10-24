import * as z from "zod";

export const baseAccountSchema = z.object({
  username: z
    .string()
    .min(3, { message: "Username must be at least 3 characters" })
    .max(50, { message: "Username cannot be longer than 50 characters" })
    .trim(),
  password: z
    .string()
    .min(4, { message: "Password must be at least 4 characters" })
    .max(50, { message: "Password cannot be longer than 50 characters" })
    .trim(),
});

export const accountSchema = baseAccountSchema
  .extend({
    confirmPassword: z.string().trim(),
  })
  .superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword != password) {
      ctx.addIssue({
        code: "custom",
        message: "The passwords do not match",
        path: ["confirmPassword"],
      });
    }
  });
