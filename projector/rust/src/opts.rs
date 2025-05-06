use clap::Parser;
use std::path::PathBuf;

#[derive(Parser, Debug)]
#[clap()]
pub struct ProjectorOptions {
    pub arguments: Vec<String>,

    #[clap(short = 'p', long = "pwd")]
    pub pwd: Option<PathBuf>,

    #[clap(short = 'c', long = "config")]
    pub config: Option<PathBuf>,
}
