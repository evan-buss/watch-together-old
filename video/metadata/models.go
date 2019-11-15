package metadata

// Formats to check file names against to deterimine if they are videos
var movieFormats = []string{
	".avi",
	".flv",
	".h264",
	".m4v",
	".mkv",
	".mov",
	".mp4",
	".mpg",
	".mpeg",
	".wmv",
}

// Metadata contains the results from running FFProbe on a single file
type Metadata struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

// Stream contains information about a single stream
type Stream struct {
	Index     int    `json:"index"`
	CodecName string `json:"codec_name"`
	Profile   string `json:"profile"`
	CodecType string `json:"codec_type"`
}

// Format contains information about a single file
type Format struct {
	FileName       string     `json:"filename"`
	NbStreams      int        `json:"nb_streams"`
	FormatName     string     `json:"format_name"`
	FormatLongName string     `json:"format_long_name"`
	Duration       string     `json:"duration"`
	Size           string     `json:"size"`
	Bitrate        string     `json:"bit_rate"`
	ProbeScore     int        `json:"probe_scrore"`
	Tags           FormatTags `json:"tags"`
}

// FormatTags gives extra info about the file format
type FormatTags struct {
	Title string `json:"title"`
}

// MovieDBInfo represents a single movie returned from the Movie Metadata API
type MovieDBInfo struct {
	RowID   int    `json:"id"`
	URL     string `json:"url"`
	Poster  string `json:"poster"`
	Rating  string `json:"rating"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
	Year    string `json:"year"`
}
