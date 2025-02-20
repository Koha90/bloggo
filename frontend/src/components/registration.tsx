import React, { useState } from "react";
import axios from "axios";

interface SignUpFormState {
  username: string;
  first_name: string;
  last_name: string;
  password: string;
  confirPassword: string;
}

const SignUpForm: React.FC = () => {
  const [formData, setFormData] = useState<SignUpFormState>({
    username: "",
    first_name: "",
    last_name: "",
    password: "",
    confirPassword: "",
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const validateForm = (): string | null => {
    if (formData.password.length < 8) {
      return "Пароль должен содержать не менее 8 символов";
    }
    if (!/^[a-zA-Z0-9_]+$/.test(formData.username)) {
      return "Имя пользователя может содержать только буквы, цифры и нижнее подчёрикивание";
    }
    if (formData.password !== formData.confirPassword) {
      return "Пароли не совпадают";
    }
    return null;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    if (isSubmitting) return;

    const validationError = validateForm();
    if (validationError) {
      setError(validationError);
      return;
    }

    setIsSubmitting(true);

    try {
      const resp = await axios.post(
        "http://localhost:8080/api/v1/signup",
        formData,
        {
          headers: { "Content-Type": "application/json" },
        },
      );

      setSuccess("User created successfully!");
      console.log(resp);

      setFormData({
        username: "",
        first_name: "",
        last_name: "",
        password: "",
        confirPassword: "",
      });
    } catch (err) {
      if (axios.isAxiosError(err) && err.response) {
        const serverError = err.response.data?.error || "Unknown server error";
        setError(`${serverError}`);
      } else {
        setError((err as Error).message);
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col m-10">
      <div className="flex flex-col">
        <label className="mb-1" htmlFor="username">
          Username:
        </label>
        <input
          type="text"
          id="username"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="username"
          required
          disabled={isSubmitting}
          className="mb-2"
        />
        <label htmlFor="first_name">Имя:</label>
        <input
          type="text"
          id="first_name"
          name="first_name"
          value={formData.first_name}
          onChange={handleChange}
          placeholder="Имя"
          disabled={isSubmitting}
          className="mb-2"
        />
        <label htmlFor="last_name">Фамилия:</label>
        <input
          type="text"
          id="last_name"
          name="last_name"
          value={formData.last_name}
          onChange={handleChange}
          placeholder="Фамилия"
          disabled={isSubmitting}
          className="mb-2"
        />
        <label htmlFor="password">Пароль:</label>
        <input
          type="password"
          id="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          placeholder="Пароль"
          required
          disabled={isSubmitting}
          className="mb-2"
        />
        <label htmlFor="confirmPassword">Подтвердить пароль:</label>
        <input
          type="password"
          id="confirPassword"
          name="confirPassword"
          value={formData.confirPassword}
          onChange={handleChange}
          placeholder="Пароль"
          required
          disabled={isSubmitting}
        />
      </div>
      <button type="submit" className="bg-sky-100" disabled={isSubmitting}>
        {isSubmitting ? "Регистрация..." : "Регистрация"}
      </button>
      {error && <p className="text-red-500">Ошибка: {error}</p>}
      {success && <p className="text-green-500">{success}</p>}
    </form>
  );
};

export default SignUpForm;
