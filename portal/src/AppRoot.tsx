import React, { useMemo } from "react";
import { Routes, Route, useParams, Navigate } from "react-router-dom";
import { ApolloProvider } from "@apollo/client";

import { makeClient } from "./graphql/adminapi/apollo";
import { useAppAndSecretConfigQuery } from "./graphql/portal/query/appAndSecretConfigQuery";
import ScreenLayout from "./ScreenLayout";
import ShowLoading from "./ShowLoading";

import ProjectRootScreen from "./graphql/portal/ProjectRootScreen";
import UsersScreen from "./graphql/adminapi/UsersScreen";
import AddUserScreen from "./graphql/adminapi/AddUserScreen";
import UserDetailsScreen from "./graphql/adminapi/UserDetailsScreen";
import AddEmailScreen from "./graphql/adminapi/AddEmailScreen";
import AddPhoneScreen from "./graphql/adminapi/AddPhoneScreen";
import AddUsernameScreen from "./graphql/adminapi/AddUsernameScreen";
import ResetPasswordScreen from "./graphql/adminapi/ResetPasswordScreen";
import AuditLogScreen from "./graphql/adminapi/AuditLogScreen";
import AuditLogEntryScreen from "./graphql/adminapi/AuditLogEntryScreen";

import AnonymousUsersConfigurationScreen from "./graphql/portal/AnonymousUsersConfigurationScreen";
import SingleSignOnConfigurationScreen from "./graphql/portal/SingleSignOnConfigurationScreen";
import PasswordPolicyConfigurationScreen from "./graphql/portal/PasswordPolicyConfigurationScreen";
import ForgotPasswordConfigurationScreen from "./graphql/portal/ForgotPasswordConfigurationScreen";
import ApplicationsConfigurationScreen from "./graphql/portal/ApplicationsConfigurationScreen";
import CreateOAuthClientScreen from "./graphql/portal/CreateOAuthClientScreen";
import EditOAuthClientScreen from "./graphql/portal/EditOAuthClientScreen";
import CustomDomainListScreen from "./graphql/portal/CustomDomainListScreen";
import VerifyDomainScreen from "./graphql/portal/VerifyDomainScreen";
import UISettingsScreen from "./graphql/portal/UISettingsScreen";
import LocalizationConfigurationScreen from "./graphql/portal/LocalizationConfigurationScreen";
import InviteAdminScreen from "./graphql/portal/InviteAdminScreen";
import PortalAdminsSettings from "./graphql/portal/PortalAdminsSettings";
import WebhookConfigurationScreen from "./graphql/portal/WebhookConfigurationScreen";
import AdminAPIConfigurationScreen from "./graphql/portal/AdminAPIConfigurationScreen";
import LoginIDConfigurationScreen from "./graphql/portal/LoginIDConfigurationScreen";
import AuthenticatorConfigurationScreen from "./graphql/portal/AuthenticatorConfigurationScreen";
import VerificationConfigurationScreen from "./graphql/portal/VerificationConfigurationScreen";
import BiometricConfigurationScreen from "./graphql/portal/BiometricConfigurationScreen";
import SubscriptionScreen from "./graphql/portal/SubscriptionScreen";
import SMTPConfigurationScreen from "./graphql/portal/SMTPConfigurationScreen";
import StandardAttributesConfigurationScreen from "./graphql/portal/StandardAttributesConfigurationScreen";
import CustomAttributesConfigurationScreen from "./graphql/portal/CustomAttributesConfigurationScreen";
import EditCustomAttributeScreen from "./graphql/portal/EditCustomAttributeScreen";
import CreateCustomAttributeScreen from "./graphql/portal/CreateCustomAttributeScreen";
import AccountDeletionConfigurationScreen from "./graphql/portal/AccountDeletionConfigurationScreen";
import AnalyticsScreen from "./graphql/portal/AnalyticsScreen";

const AppRoot: React.FC = function AppRoot() {
  const { appID } = useParams();
  const client = useMemo(() => {
    return makeClient(appID);
  }, [appID]);

  // NOTE: check if appID actually exist in authorized app list
  const { effectiveAppConfig, loading, error } =
    useAppAndSecretConfigQuery(appID);
  if (loading) {
    return <ShowLoading />;
  }

  // if node is null after loading without error, treat this as invalid
  // request, frontend cannot distinguish between inaccessible and not found
  const isInvalidAppID = error == null && effectiveAppConfig == null;

  // redirect to app list if app id is invalid
  if (isInvalidAppID) {
    return <Navigate to="/projects" replace={true} />;
  }

  return (
    <ApolloProvider client={client}>
      <ScreenLayout>
        <Routes>
          <Route path="/" element={<ProjectRootScreen />} />
          <Route path="/analytics" element={<AnalyticsScreen />} />
          <Route path="/users/" element={<UsersScreen />} />
          <Route path="/users/add-user/" element={<AddUserScreen />} />
          <Route
            path="/users/:userID/"
            element={<Navigate to="details/" replace={true} />}
          />
          <Route
            path="/users/:userID/details/"
            element={<UserDetailsScreen />}
          />
          <Route
            path="/users/:userID/details/add-email"
            element={<AddEmailScreen />}
          />
          <Route
            path="/users/:userID/details/add-phone"
            element={<AddPhoneScreen />}
          />
          <Route
            path="/users/:userID/details/add-username"
            element={<AddUsernameScreen />}
          />
          <Route
            path="/users/:userID/details/reset-password"
            element={<ResetPasswordScreen />}
          />
          <Route
            path="/configuration/authentication/login-id"
            element={<LoginIDConfigurationScreen />}
          />
          <Route
            path="/configuration/authentication/authenticators"
            element={<AuthenticatorConfigurationScreen />}
          />
          <Route
            path="/configuration/authentication/verification"
            element={<VerificationConfigurationScreen />}
          />
          <Route
            path="/configuration/anonymous-users"
            element={<AnonymousUsersConfigurationScreen />}
          />
          <Route
            path="/configuration/biometric"
            element={<BiometricConfigurationScreen />}
          />
          <Route
            path="/configuration/single-sign-on"
            element={<SingleSignOnConfigurationScreen />}
          />
          <Route
            path="/configuration/password-policy"
            element={<PasswordPolicyConfigurationScreen />}
          />
          <Route
            path="/advanced/password-reset-code"
            element={<ForgotPasswordConfigurationScreen />}
          />
          <Route
            path="/configuration/apps"
            element={<ApplicationsConfigurationScreen />}
          />
          <Route
            path="/configuration/apps/add"
            element={<CreateOAuthClientScreen />}
          />
          <Route
            path="/configuration/apps/:clientID/edit"
            element={<EditOAuthClientScreen />}
          />
          <Route path="/custom-domains" element={<CustomDomainListScreen />} />
          <Route
            path="/custom-domains/:domainID/verify"
            element={<VerifyDomainScreen />}
          />
          <Route
            path="/configuration/ui-settings"
            element={<UISettingsScreen />}
          />
          <Route
            path="/configuration/localization"
            element={<LocalizationConfigurationScreen />}
          />
          <Route
            path="/configuration/user-profile/standard-attributes"
            element={<StandardAttributesConfigurationScreen />}
          />
          <Route
            path="/configuration/user-profile/custom-attributes"
            element={<CustomAttributesConfigurationScreen />}
          />
          <Route
            path="/configuration/user-profile/custom-attributes/:index/edit"
            element={<EditCustomAttributeScreen />}
          />
          <Route
            path="/configuration/user-profile/custom-attributes/add"
            element={<CreateCustomAttributeScreen />}
          />
          <Route
            path="/configuration/smtp"
            element={<SMTPConfigurationScreen />}
          />
          <Route path="/portal-admins" element={<PortalAdminsSettings />} />
          <Route path="/portal-admins/invite" element={<InviteAdminScreen />} />
          <Route
            path="/advanced/webhooks"
            element={<WebhookConfigurationScreen />}
          />
          <Route path="/billing" element={<SubscriptionScreen />} />
          <Route
            path="/advanced/admin-api"
            element={<AdminAPIConfigurationScreen />}
          />
          <Route
            path="/advanced/account-deletion"
            element={<AccountDeletionConfigurationScreen />}
          />
          <Route path="/audit-log/" element={<AuditLogScreen />} />
          <Route
            path="/audit-log/:logID/"
            element={<Navigate to="details/" replace={true} />}
          />
          <Route
            path="/audit-log/:logID/details/"
            element={<AuditLogEntryScreen />}
          />
        </Routes>
      </ScreenLayout>
    </ApolloProvider>
  );
};

export default AppRoot;
