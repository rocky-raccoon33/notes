---
title: Spring Web MVC
comments: true
description: 表现层框架 MVC 是一种软件架构思想，是一种开发模式，将软件划分为三种不同类型的模块，分别是模型，视图，和控制器
date: 2022-02-01
category: Web Development
section: blog
tags:
    - Spring
---

![](./img/filters_vs_interceptors.jpeg)

## `1` Filters

- `Filters` are part of the web server and not the Spring framework, it intercep requests before they reach the `DispatcherServlet`, making them ideal for coarse-grained tasks such as:

- Authentication
- Logging and auditing
- Image and data compression
- Any functionality we want to be decoupled from Spring MVC

```java

@Component
public class LogFilter implements Filter {

    private Logger logger = LoggerFactory.getLogger(LogFilter.class);

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain) 
      throws IOException, ServletException {
        logger.info("Hello from: " + request.getLocalAddr());
        chain.doFilter(request, response);
    }
}

```

## `2` HandlerInterceptors

- `HandlerInterceptors` are part of the Spring MVC framework and sit between the DispatcherServlet and the controllers.
it intercepts requests between the `DispatcherServlet` and the `Controllers`. This is done within the Spring MVC framework, providing access to the Handler and ModelAndView objects. This reduces duplication and allows for more fine-grained functionality such as:

- Handling cross-cutting concerns such as application logging
- Detailed authorization checks
- Manipulating the Spring context or model

```java

public class LogInterceptor implements HandlerInterceptor {

    private Logger logger = LoggerFactory.getLogger(LogInterceptor.class);

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) 
      throws Exception {
        logger.info("preHandle");
        return true;
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView modelAndView) 
      throws Exception {
        logger.info("postHandle");
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) 
      throws Exception {
        logger.info("afterCompletion");
    }

}

```
