use anyhow::Result;
use clap::Parser;
use projector::{config::Config, opts::ProjectorOptions};

fn main() -> Result<()> {
    let options: Config = ProjectorOptions::parse().try_into()?;
    println!("{:?}", options);

    return Ok(());
}
