import React from "react";
import { Link as MuiLink, LinkProps as MuiLinkProps } from "@mui/material";
import { useLocale } from "@/hooks";

export const Link = React.forwardRef<HTMLAnchorElement, MuiLinkProps>(
  function Link({ children, href, sx, ...props }, ref) {
    const { locale } = useLocale();

    return (
      <MuiLink
        ref={ref}
        href={`/${locale}${href}`}
        sx={{
          color: "inherit",
          textDecoration: "none",
          ...sx,
        }}
        {...props}
      >
        {children}
      </MuiLink>
    );
  },
);
