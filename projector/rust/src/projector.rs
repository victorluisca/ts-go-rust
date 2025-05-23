use std::{collections::HashMap, path::PathBuf};

use anyhow::Result;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct ProjectorData {
    pub projector: HashMap<PathBuf, HashMap<String, String>>,
}

pub struct Projector {
    config: PathBuf,
    pwd: PathBuf,
    data: ProjectorData,
}

fn default_data() -> ProjectorData {
    return ProjectorData {
        projector: HashMap::new(),
    };
}

impl Projector {
    pub fn get_value_all(&self) -> HashMap<&String, &String> {
        let mut current = Some(self.pwd.as_path());
        let mut paths = vec![];
        let mut output = HashMap::new();

        while let Some(path) = current {
            paths.push(path);
            current = path.parent();
        }

        for path in paths.into_iter().rev() {
            if let Some(map) = self.data.projector.get(path) {
                output.extend(map.iter());
            }
        }

        return output;
    }

    pub fn get_value(&self, key: &str) -> Option<&String> {
        let mut current = Some(self.pwd.as_path());
        let mut output = None;

        while let Some(path) = current {
            if let Some(dir) = self.data.projector.get(path) {
                if let Some(value) = dir.get(key) {
                    output = Some(value);
                    break;
                }
            }
            current = path.parent()
        }

        return output;
    }

    pub fn set_value(&mut self, key: String, value: String) {
        self.data
            .projector
            .entry(self.pwd.clone())
            .or_default()
            .insert(key, value);
    }

    pub fn delete_value(&mut self, key: &str) {
        self.data
            .projector
            .get_mut(&self.pwd)
            .map(|x| x.remove(key));
    }

    pub fn save(&self) -> Result<()> {
        if let Some(p) = self.config.parent() {
            if !std::fs::metadata(&p).is_ok() {
                std::fs::create_dir_all(p)?;
            }
        }

        let contents = serde_json::to_string(&self.data)?;
        std::fs::write(&self.config, contents)?;

        return Ok(());
    }

    pub fn from_config(config: PathBuf, pwd: PathBuf) -> Self {
        if std::fs::metadata(&config).is_ok() {
            let contents = std::fs::read_to_string(&config);
            let contents = contents.unwrap_or(String::from("{\"projector\":{}}"));
            let data = serde_json::from_str(&contents);
            let data = data.unwrap_or(default_data());

            return Projector { config, pwd, data };
        }
        return Projector {
            config,
            pwd,
            data: default_data(),
        };
    }
}

#[cfg(test)]
mod test {
    use std::{collections::HashMap, path::PathBuf};

    use collection_macros::hashmap;

    use super::Projector;

    fn get_data() -> HashMap<PathBuf, HashMap<String, String>> {
        return hashmap! {
            PathBuf::from("/") => hashmap! {
                "foo".into() => "bar1".into(),
                "bar".into() => "baz".into()
            },
            PathBuf::from("/foo") => hashmap! {
                "foo".into() => "bar2".into()
            },
            PathBuf::from("/foo/bar") => hashmap! {
                "foo".into() => "bar3".into()
            }
        };
    }

    fn get_projector(pwd: PathBuf) -> Projector {
        return Projector {
            config: PathBuf::from(""),
            pwd,
            data: super::ProjectorData {
                projector: get_data(),
            },
        };
    }

    #[test]
    fn get_value() {
        let projector = get_projector(PathBuf::from("/foo/bar"));

        assert_eq!(projector.get_value("foo"), Some(&String::from("bar3")));
        assert_eq!(projector.get_value("bar"), Some(&String::from("baz")));
    }

    #[test]
    fn set_value() {
        let mut projector = get_projector(PathBuf::from("/foo/bar"));

        projector.set_value(String::from("foo"), String::from("bar4"));
        projector.set_value(String::from("bar"), String::from("foo"));

        assert_eq!(projector.get_value("foo"), Some(&String::from("bar4")));
        assert_eq!(projector.get_value("bar"), Some(&String::from("foo")));
    }

    #[test]
    fn delete_value() {
        let mut projector = get_projector(PathBuf::from("/foo/bar"));

        projector.delete_value("foo");
        projector.delete_value("bar");

        assert_eq!(projector.get_value("foo"), Some(&String::from("bar2")));
        assert_eq!(projector.get_value("bar"), Some(&String::from("baz")));
    }
}
