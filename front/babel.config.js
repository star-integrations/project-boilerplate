module.exports = {
  presets: [
    [
      '@babel/preset-env',
      {
        useBuiltIns: 'usage',
        targets: '> 0.25%, not dead',
        corejs: { version: '3.8', proposals: true },
      },
    ],
  ],
  env: {
    test: {
      presets: [['@babel/preset-env', { targets: { node: 'current' } }]],
    },
  },
  plugins: [
    ['@babel/plugin-proposal-private-property-in-object', { loose: true }],
  ],
};
