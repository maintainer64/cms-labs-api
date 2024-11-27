import React from 'react';
import { AuthLayoutWrapper } from '@/components/pages/auth/authLayout';
import '@/styles/globals.css';
import { Login } from '@/components/pages/auth/login';
import AuthError from '@/components/pages/auth/error';

export default function LoginPage() {
  return (
    <AuthLayoutWrapper>
      <Login />
    </AuthLayoutWrapper>
  );
}

export function LoginError() {
  return (
    <AuthLayoutWrapper>
      <AuthError />
    </AuthLayoutWrapper>
  );
}
