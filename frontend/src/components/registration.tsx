import React, { useState } from "react";
import axios from "axios";

interface SignUpFormState {
  username: string;
  first_name: string;
  last_name: string;
  password: string;
}

const SignUpForm: React.FC = () => {
  const [formData, setFormData] = useState<SignUpFormState>({
    username: "",
    first_name: "",
    last_name: "",
    password: "",
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    // Проверка длины пароля
    if (formData.password.length < 8) {
      setError("Пароль должен содержать не менее 8 символов");
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
      });
    } catch (err) {
      if (axios.isAxiosError(err) && err.response) {
        setError(
          `Server error: ${err.response?.statusText || "Unknown Error"} (${err.response.status})`,
        );
      } else {
        setError((err as Error).message);
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label className="ml-2" htmlFor="username">
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
        />
        <label htmlFor="first_name">Имя:</label>
        <input
          type="text"
          id="first_name"
          name="first_name"
          value={formData.first_name}
          onChange={handleChange}
          placeholder="Имя"
        />
        <label htmlFor="last_name">Фамилия:</label>
        <input
          type="text"
          id="last_name"
          name="last_name"
          value={formData.last_name}
          onChange={handleChange}
          placeholder="Фамилия"
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
        />
      </div>
      <button type="submit" disabled={isSubmitting}>
        {isSubmitting ? "Регистрация..." : "Регистрация"}
      </button>
      {error && <p className="text-red-500">Ошибка: {error}</p>}
      {success && <p className="text-green-500">{success}</p>}
    </form>
  );
};

export default SignUpForm;
