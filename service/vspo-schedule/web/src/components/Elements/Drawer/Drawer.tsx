import {
  Badge,
  Box,
  Chip,
  Divider,
  Link,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Stack,
  SwipeableDrawer,
  Typography,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import React from "react";
import { faTwitch } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  getNavigationRouteInfo,
  NavigationRouteId,
} from "@/constants/navigation";
import { DrawerIcon } from "../Icon";
import { ThemeToggleButton } from "../Button";
import { useTranslation } from "next-i18next";
import { TimeZoneSelector } from "../Control";
import { useTimeZoneContext } from "@/hooks";
import { DEFAULT_LOCALE } from "@/lib/Const";
import { useRouter } from "next/router";

const drawerNavigationSections: NavSectionProps[] = [
  {
    heading: "Main Section",
    links: [
      { id: "live" },
      { id: "upcoming" },
      { id: "archive" },
      { id: "freechat" },
    ],
  },
  {
    heading: "Clips Section",
    links: [
      { id: "clip", isBeta: true },
      {
        id: "twitch-clip",
        isBeta: true,
        supplementaryIcon: <FontAwesomeIcon icon={faTwitch} />,
      },
    ],
  },
  {
    heading: "Help Section",
    links: [
      { id: "about" },
      { id: "site-news" },
      { id: "qa" },
      { id: "discord" },
    ],
  },
];

const StyledListItemIcon = styled(ListItemIcon)(() => ({
  minWidth: "32px",
}));

const StyledChip = styled(Chip)(({ theme }) => ({
  backgroundColor: "transparent",
  border: "1px solid",
  borderColor: theme.vars.palette.customColors.darkBlue,
  color: theme.vars.palette.customColors.darkBlue,

  [theme.getColorSchemeSelector("dark")]: {
    borderColor: "white",
    color: "white",
  },
}));

const StyledBadge = styled(Badge)(({ theme }) => ({
  "& .MuiBadge-badge": {
    backgroundColor: theme.vars.palette.customColors.darkBlue,
    color: "white",
    transform: "scale(0.8)",
    fontSize: "0.65em",
    right: "-28px",
    top: "1px",
  },
}));

type NavLinkProps = {
  id: NavigationRouteId;
  isBeta?: boolean;
  supplementaryIcon?: React.ReactNode;
};

type NavSectionProps = {
  heading: string;
  links: NavLinkProps[];
};

const NavSectionHeading: React.FC<{ text: string }> = ({ text }) => (
  <Typography variant="subtitle2" sx={{ padding: "8px 16px" }}>
    <StyledChip label={text} size="small" />
  </Typography>
);

const NavLink: React.FC<NavLinkProps> = ({ id, isBeta, supplementaryIcon }) => {
  const { t } = useTranslation("common");
  const { timeZone } = useTimeZoneContext();
  const router = useRouter();
  const locale = router?.locale ?? DEFAULT_LOCALE;
  const { link, isExternalLink } = getNavigationRouteInfo(id, timeZone, locale);
  const buttonProps = isExternalLink
    ? { component: "a", target: "_blank", rel: "noopener noreferrer" }
    : { component: Link };
  const label = t(`drawer.pages.${id}`);

  return (
    <ListItemButton
      href={link}
      {...buttonProps}
      disableRipple
      sx={{
        "&:hover": {
          backgroundColor: "transparent",
        },
      }}
    >
      <StyledListItemIcon>
        <DrawerIcon id={id} />
      </StyledListItemIcon>
      {supplementaryIcon && (
        <Box sx={{ marginRight: "4px" }}>{supplementaryIcon}</Box>
      )}
      {isBeta ? (
        <StyledBadge badgeContent="Beta">
          <ListItemText primary={label} />
        </StyledBadge>
      ) : (
        <ListItemText primary={label} />
      )}
    </ListItemButton>
  );
};

const DrawerNavigation: React.FC<{
  sections: NavSectionProps[];
}> = ({ sections }) => {
  return (
    <nav>
      <List>
        {sections.map((sectionProps, index) => (
          <React.Fragment key={index}>
            {index !== 0 && <Divider component="li" />}
            <ListItem
              disablePadding
              sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "stretch",
              }}
            >
              <NavSectionHeading text={sectionProps.heading} />
              <List disablePadding>
                {sectionProps.links.map((linkProps) => (
                  <ListItem key={linkProps.id} disablePadding>
                    <NavLink {...linkProps} />
                  </ListItem>
                ))}
              </List>
            </ListItem>
          </React.Fragment>
        ))}
      </List>
    </nav>
  );
};

const AppBarOffset = styled("div")(({ theme }) => theme.mixins.toolbar);

type DrawerProps = {
  open: boolean;
  onOpen: () => void;
  onClose: () => void;
};

export const CustomDrawer: React.FC<DrawerProps> = ({
  open,
  onOpen,
  onClose,
}) => {
  return (
    <SwipeableDrawer
      anchor="left"
      open={open}
      onOpen={onOpen}
      onClose={onClose}
      PaperProps={{
        style: {
          width: "240px",
          zIndex: 1200,
        },
      }}
      ModalProps={{
        keepMounted: true,
        disablePortal: true,
      }}
      SwipeAreaProps={{
        style: {
          position: "absolute",
        },
      }}
    >
      <AppBarOffset />

      <Stack direction="column" justifyContent="space-between" flex="1">
        <DrawerNavigation sections={drawerNavigationSections} />
        <Divider />
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: "16px",
            padding: "20px 12px",
          }}
        >
          <TimeZoneSelector />
          <ThemeToggleButton />
        </Box>
      </Stack>
    </SwipeableDrawer>
  );
};
