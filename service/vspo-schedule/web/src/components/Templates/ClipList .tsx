import React, { useContext, useState } from "react";
import {
  Container,
  Grid,
  Card,
  CardActionArea,
  CardContent,
  Typography,
  Box,
  Pagination,
  Avatar,
  Chip,
  PaletteColor,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import { Clip, Platform } from "@/types/streaming";
import { isTrending } from "@/lib/utils";
import { members } from "@/data/members";
import { PlayArrow } from "@mui/icons-material";
import { EmbedModeContext } from "@/context/EmbedMode";
import { useModal } from "@/hooks";
import dynamic from "next/dynamic";
import Image from "next/image";

const LivestreamDetailsModal = dynamic(
  () => import("../Elements/Modal").then((mod) => mod.LivestreamDetailsModal),
  { ssr: false },
);

type Props = {
  clips: Clip[];
};

const StyledCardMedia = styled(Box)({
  paddingTop: "56.25%", // 16:9 aspect ratio
  objectFit: "contain",
});

const StyledAvatar = styled(Avatar)(({ theme }) => ({
  width: 48,
  height: 48,
  [theme.breakpoints.down("sm")]: {
    width: 32,
    height: 32,
  },
}));

const formatViewCount = (viewCount: number) => {
  return Math.floor(viewCount / 100_000) * 10;
};

const getClipLabel = (clip: Clip) => {
  if (isTrending(clip)) {
    return { label: "急上昇", color: "error" };
  }
  // if (isPopular(clip)) {
  //   return { label: "人気", color: "warning" };
  // }
  // if (isNew(clip)) {
  //   return { label: "新着", color: "success" };
  // }
  return null;
};

export const ClipList: React.FC<Props> = ({ clips }) => {
  const [page, setPage] = useState(1);
  const [isEmbedMode] = useContext(EmbedModeContext);
  const { isOpen, openModal: baseOpenModal, closeModal } = useModal();
  const [clickedClip, setClickedClip] = useState<Clip | null>(null);
  const clipsPerPage = 24;

  const handleChange = (event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };
  const openModal = (clip: Clip) => {
    setClickedClip(clip);
    baseOpenModal();
  };

  const paginatedClips = clips.slice(
    (page - 1) * clipsPerPage,
    page * clipsPerPage,
  );

  return (
    <Container maxWidth="lg" sx={{ paddingTop: "50px" }}>
      <Grid container spacing={3}>
        {paginatedClips.map((clip) => {
          const iconUrl =
            clip.platform === Platform.Twitch
              ? members
                  .filter((m) => m.twitchChannelId === clip.channelId)
                  .at(0)?.iconUrl
              : clip.iconUrl;

          const clipLabel = getClipLabel(clip);
          return (
            <Grid item xs={12} sm={6} md={4} key={clip.id}>
              <Box
                height="100%"
                display="flex"
                sx={{
                  position: "relative",
                  ...(clipLabel ? { border: "3px solid red" } : {}),
                }}
              >
                {clipLabel && (
                  <Chip
                    label={clipLabel.label}
                    sx={{
                      position: "absolute",
                      transform: "translateY(-12px)",
                      zIndex: "3",
                      top: "0px",
                      right: "8px",
                      height: "24px",
                      backgroundColor: (theme) =>
                        (
                          theme.vars.palette[
                            clipLabel.color as keyof typeof theme.vars.palette
                          ] as PaletteColor
                        ).main,
                      color: "white",
                    }}
                  />
                )}
                <Card
                  sx={{
                    height: "100%",
                    display: "flex",
                    flexDirection: "column",
                    width: "100%",
                    position: "relative",
                  }}
                >
                  <CardActionArea onClick={() => openModal(clip)}>
                    <StyledCardMedia position="relative">
                      <Image
                        src={clip.thumbnailUrl}
                        alt={clip.title}
                        fill
                        style={{ objectFit: "cover" }}
                      />
                    </StyledCardMedia>
                    <CardContent
                      sx={{
                        flexGrow: 1,
                        display: "flex",
                        flexDirection: "column",
                      }}
                    >
                      <Typography
                        gutterBottom
                        variant="body1"
                        component="div"
                        sx={{
                          maxHeight: "3em",
                          minHeight: "3em",
                          display: "-webkit-box",
                          WebkitLineClamp: 2,
                          WebkitBoxOrient: "vertical",
                          overflow: "hidden",
                        }}
                      >
                        {clip.title}
                      </Typography>
                      <Box
                        mt="auto"
                        sx={{
                          display: "flex",
                          alignItems: "center",
                          gap: "0.5rem",
                          justifyContent: "space-between",
                        }}
                      >
                        <Box
                          mt="auto"
                          sx={{
                            display: "flex",
                            alignItems: "center",
                            gap: "0.5rem",
                          }}
                        >
                          <Typography
                            variant="body2"
                            color="textSecondary"
                            component="p"
                          >
                            {clip.channelTitle}
                          </Typography>
                          <StyledAvatar>
                            {iconUrl && (
                              <Image
                                src={iconUrl}
                                alt={clip.channelTitle}
                                fill
                                style={{ objectFit: "cover" }}
                              />
                            )}
                          </StyledAvatar>
                        </Box>
                        {clip.viewCount && Number(clip.viewCount) > 100000 && (
                          <Chip
                            icon={
                              <PlayArrow
                                fontSize="small"
                                sx={{ color: "white !important" }}
                              />
                            }
                            label={`${formatViewCount(
                              Number(clip.viewCount),
                            )}万以上`}
                            sx={{
                              height: "24px",
                              width: "90px",
                              backgroundColor: "rgba(0, 0, 0, 0.6)",
                              color: "white",
                              fontSize: "10px",
                            }}
                          />
                        )}
                      </Box>
                    </CardContent>
                  </CardActionArea>
                </Card>
              </Box>
            </Grid>
          );
        })}
      </Grid>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          padding: "1rem",
        }}
      >
        <Pagination
          count={Math.ceil(clips.length / clipsPerPage)}
          page={page}
          onChange={handleChange}
          color="primary"
          size="medium"
          showFirstButton
          showLastButton
        />
      </Box>
      {clickedClip && (
        <LivestreamDetailsModal
          key={clickedClip.id}
          clip={clickedClip}
          open={isOpen}
          onClose={() => {
            setClickedClip(null);
            closeModal();
          }}
          isDefaultEmbedMode={isEmbedMode}
        />
      )}
    </Container>
  );
};
