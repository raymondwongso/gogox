# gogox

Gogox is a collection of reusable golang components

Below listed available components, they have their own README section (linked):
1. [cache](/cache/README.md) is a cache adapter, so that your application doesn't need to know the backend implementation.
2. [errorx](/errorx/README.md) is a custom error handler components. It contains more robust attributes and works with grpc-gateway (must use gogox grpc components)
3. [grpc](/grpc/README.md) is a GRPC specific components, such as log and trace interceptors.
4. [http](/http/README.md) is a HTTP specific components, such as middlwares.
5. [log](/log/README.md) is a log adapter, so that your application doesn't need to know the logger implementation.
6. [stats](/stats/README.md) is a stats adapter, so that your application doesn't need to know the stats implementation.
7. [trace](/trace/README.md) is a trace ID generator.
8. [sugar](/sugar/README.md) is a sugar syntax package.