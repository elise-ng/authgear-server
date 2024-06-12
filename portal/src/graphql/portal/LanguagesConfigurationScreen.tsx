import React from "react";
import { useParams } from "react-router-dom";
import { produce } from "immer";
import cn from "classnames";
import { PortalAPIAppConfig } from "../../types";
import { clearEmptyObject } from "../../util/misc";
import { useAppConfigForm } from "../../hook/useAppConfigForm";
import FormContainer from "../../FormContainer";
import ScreenContent from "../../ScreenContent";
import ScreenTitle from "../../ScreenTitle";
import { FormattedMessage } from "@oursky/react-messageformat";

interface ConfigFormState {
  supportedLanguages: string[];
  fallbackLanguage: string;
}

function constructFormState(config: PortalAPIAppConfig): ConfigFormState {
  const fallbackLanguage = config.localization?.fallback_language ?? "en";
  return {
    fallbackLanguage,
    supportedLanguages: config.localization?.supported_languages ?? [
      fallbackLanguage,
    ],
  };
}

function constructConfig(
  config: PortalAPIAppConfig,
  _initialState: ConfigFormState,
  currentState: ConfigFormState
): PortalAPIAppConfig {
  return produce(config, (config) => {
    config.localization = config.localization ?? {};
    config.localization.fallback_language = currentState.fallbackLanguage;
    config.localization.supported_languages = currentState.supportedLanguages;
    clearEmptyObject(config);
  });
}

const LanguagesConfigurationScreen: React.VFC =
  function LanguagesConfigurationScreen() {
    const { appID } = useParams() as { appID: string };
    const appConfigForm = useAppConfigForm({
      appID,
      constructFormState,
      constructConfig,
    });
    return (
      <FormContainer form={appConfigForm} canSave={true}>
        <ScreenContent>
          <ScreenTitle className={cn("col-span-8", "tablet:col-span-full")}>
            <FormattedMessage id="LanguagesConfigurationScreen.title" />
          </ScreenTitle>
        </ScreenContent>
      </FormContainer>
    );
  };

export default LanguagesConfigurationScreen;
