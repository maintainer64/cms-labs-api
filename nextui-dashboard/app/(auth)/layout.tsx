import React from 'react';
import { AuthLayoutWrapper } from '@/components/pages/auth/authLayout';
import '@/styles/globals.css';
import { Login } from '@/components/pages/auth/login';
import AuthError from '@/components/pages/auth/error';
import { LTIAttemptSSO } from '@/components/pages/lti-attempts/create/lti-attempt-sso';

export default function LoginPage() {
  return (
    <AuthLayoutWrapper>
      <LTIAttemptSSO>
        <Login />
      </LTIAttemptSSO>
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
