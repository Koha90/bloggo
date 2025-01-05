import React, { useState } from "react";
import axios from "axios";

interface SignInFormState {
  username: string;
  password: string;
}

const SignInForm: React.FC = () => {
  const [formData, setFormData] = useState<SignInFormState>({
    username: "",
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

    if (isSubmitting) return;

    setIsSubmitting(true);

    try {
      const resp = await axios.post(
        "http://localhost:8080/api/v1/signin",
        formData,
        {
          headers: { "Content-Type": "application/json" },
        },
      );

      setSuccess("Loged successfully!");
      console.log(resp);

      setFormData({
        username: "",
        password: "",
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
    <form onSubmit={handleSubmit} className="m-10 flex flex-col">
      <div className="m-0 flex flex-col">
        <label htmlFor="username" className="mb-1">
          Username:
        </label>
        <input
          type="text"
          id="username"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Username"
          required
          disabled={isSubmitting}
          className="mb-2"
        />
        <label htmlFor="password" className="mb-1">
          Пароль:
        </label>
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
      </div>
      <button
        type="submit"
        className="bg-indigo-50 flex-none"
        disabled={isSubmitting}
      >
        {isSubmitting ? "Вход..." : "Вход"}
      </button>
      {error && <p className="text-red-500">Ошибка: {error}</p>}
      {success && <p className="text-green-500">{success}</p>}
    </form>
  );
};

export default SignInForm;
