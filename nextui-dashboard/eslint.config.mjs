import globals from 'globals';
import pluginJs from '@eslint/js';
import tseslint from 'typescript-eslint';
import pluginReact from 'eslint-plugin-react';
import pluginPrettier from 'eslint-config-prettier';


/** @type {import('eslint').Linter.Config[]} */
export default [
  { files: ['**/*.{js,mjs,cjs,ts,jsx,tsx}'] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
  pluginPrettier,
  {
    'rules': {
      '@typescript-eslint/ban-ts-comment': [
        'off'
      ],
      'no-unused-vars': 'off',
      '@typescript-eslint/no-unused-vars': ['off'],
      '@typescript-eslint/no-explicit-any': 'off',
      'no-prototype-builtins': 'off',
      '@typescript-eslint/no-unused-expressions': 'off',
      // suppress errors for missing 'import React' in files
      // allow jsx syntax in js files (for next.js project)
      'react/jsx-filename-extension': [
        1,
        {
          'extensions': [
            '.js',
            '.jsx',
            '.ts',
            '.tsx'
          ]
        }
      ],
      'class-methods-use-this': 'off',
      'react-hooks/rules-of-hooks': 'off',
      'react/no-array-index-key': 'off',
      'max-classes-per-file': [
        'error',
        {
          'ignoreExpressions': true,
          'max': 2
        }
      ],
      'react/require-default-props': 'off',
      'react/react-in-jsx-scope': [
        'off'
      ],
      'react/jsx-uses-react': [
        'off'
      ],
      'react/jsx-props-no-spreading': [
        'off'
      ],
      'react/no-unescaped-entities': [
        'off'
      ]
    }
  }
];